/**
 *  通过本文，你将可以做到：

	创建自己的区块链

	理解 hash 函数是如何保持区块链的完整性

	如何创造并添加新的块

	多个节点如何竞争生成块

	通过浏览器来查看整个链

	所有其他关于区块链的基础知识

	但是，对于比如工作量证明算法（PoW）以及权益证明算法（PoS）这类的共识算法文章中将不会涉及。同时为了让你更清楚得查看区块链以及块的添加，我们将网络交互的过程简化了，关于 P2P 网络比如“全网广播”这个过程等内容将在下一篇文章中补上。
*/
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

/**
 *  Index 是这个块在整个链中的位置

	Timestamp 显而易见就是块生成时的时间戳

	Hash 是这个块通过 SHA256 算法生成的散列值

	PrevHash 代表前一个块的 SHA256 散列值

	BPM 每分钟心跳数，也就是心率
*/
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

// 我们使用散列算法（SHA256）来确定和维护链中块和块正确的顺序，确保每一个块的 PrevHash 值等于前一个块中的 Hash 值，这样就以正确的块顺序构建出链
var Blockchain []Block

/**	我们为什么需要散列？主要是两个原因：

在节省空间的前提下去唯一标识数据。散列是用整个块的数据计算得出，在我们的例子中，将整个块的数据通过 SHA256 计算成一个定长不可伪造的字符串。

维持链的完整性。通过存储前一个块的散列值，我们就能够确保每个块在链中的正确顺序。任何对数据的篡改都将改变散列值，同时也就破坏了链。以我们从事的医疗健康领域为例，比如有一个恶意的第三方为了调整“人寿险”的价格，而修改了一个或若干个块中的代表不健康的 BPM 值，那么整个链都变得不可信了。
*/
func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

/**
这个 calculateHash 函数接受一个块，通过块中的 Index，Timestamp，BPM，以及 PrevHash 值来计算出 SHA256 散列值。接下来我们就能便携一个生成块的函数：
*/
func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
	/**
	其中，Index 是从给定的前一块的 Index 递增得出，时间戳是直接通过 time.Now() 函数来获得的，Hash 值通过前面的 calculateHash 函数计算得出，PrevHash 则是给定的前一个块的 Hash 值。
	*/
}

/**
搞定了块的生成，接下来我们需要有函数帮我们判断一个块是否有被篡改。检查 Index 来看这个块是否正确得递增，检查 PrevHash 与前一个块的 Hash 是否一致，再来通过 calculateHash 检查当前块的 Hash 值是否正确。通过这几步我们就能写出一个校验函数
*/
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

/**
除了校验块以外，我们还会遇到一个问题：两个节点都生成块并添加到各自的链上，那我们应该以谁为准？这里的细节我们留到下一篇文章，这里先让我们记住一个原则：始终选择最长的链。
*/
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

/****************************传统的web服务来查看整个区块链******************************/

func run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on ", os.Getenv("ADDR"))
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
