// Client code
package main

func main() {
    lfu := new(lfu)
    cache := initCache(lfu)

    cache.add("a", "1")
    cache.add("b", "2")

    cache.add("c", "3")

    lru := new(lru)
    cache.setEvictionAlgo(lru)

    cache.add("d", "4")

    fifo := new(fifo)
    cache.setEvictionAlgo(fifo)

    cache.add("e", "5")

}