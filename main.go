
package main

import (
	"github.com/9thchain/trie/db"
	"github.com/9thchain/trie/common"
	"github.com/9thchain/trie/trie"
	"fmt"
)

func main() {

	diskdb, _ := db.NewMemDatabase()
	trie, _ := trie.New(common.Hash{}, trie.NewDatabase(diskdb))

	trie.Update([]byte("a"), []byte("a-test-string"))
	trie.Update([]byte("b"), []byte("b-test-string"))

	root := trie.Hash()
	fmt.Println("root hash:", common.Bytes2Hex(root.Bytes()))
}
