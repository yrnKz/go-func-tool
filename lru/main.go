package main

type LRUCache struct {
	size   int
	p      map[int]*list
	head   int
	last   int
	length int
}

type list struct {
	next  *list
	prior *list
	value int
}

func Constructor(capacity int) LRUCache {
	lRUCache := LRUCache{
		size:   capacity,
		p:      make(map[int]*list, capacity),
		head:   -1,
		last:   -1,
		length: 0,
	}

	return lRUCache
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.p[key]
	if !ok {
		return -1
	}
	if key != this.head {
		v.prior.next = v.next
		if v.next != nil {
			v.next.prior = v.prior
		} else {
			this.last = v.prior.value
		}
		v.prior = nil
		v.next = this.p[this.head]
		this.head = v.value
	}
	return v.value
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.p[key]; ok {
		v.value = value
		return
	}
	if this.length < this.size {
		node := list{
			next:  nil,
			prior: nil,
			value: value,
		}
		if this.head == -1 {
			this.head = value
			this.last = value
		} else {
			node.next = this.p[this.head]
			this.p[this.head].prior = &node
			this.head = value
		}
		this.p[key] = &node
		this.length++
	} else {
		node := list{
			next:  nil,
			prior: nil,
			value: value,
		}
		node.next = this.p[this.head]
		this.p[this.head].prior = &node
		this.head = value
		this.p[key] = &node
		tempV := this.p[this.last].prior.value
		this.p[this.last].prior.next = nil
		this.p[this.last].prior = nil
		delete(this.p, this.last)
		this.last = tempV
	}
}

func main() {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Get(1)
	l.Put(3, 3)
	l.Get(2)
}
