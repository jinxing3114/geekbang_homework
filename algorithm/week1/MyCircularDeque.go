package main

type MyCircularDeque struct {
	de []int
	head, last, lastInd int
}


func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		de: make([]int, k),
		lastInd: k-1,
		head:-1,
		last:-1,
	}
}


func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.head == -1 {
		this.de[0] = value
		this.head = 0
		this.last = 0
		return true
	}
	var nextInd = this.lastInd
	if this.head != 0 {
		nextInd = this.head-1
	}

	if this.last == nextInd {
		return false
	} else {
		this.head = nextInd
		this.de[nextInd] = value
		return true
	}
}


func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.head == -1 {
		this.de[0] = value
		this.head = 0
		this.last = 0
		return true
	}

	var nextInd int
	if this.last != this.lastInd {
		nextInd = this.last+1
	}

	if this.head == nextInd {
		return false
	} else {
		this.last = nextInd
		this.de[nextInd] = value
		return true
	}
}


func (this *MyCircularDeque) DeleteFront() bool {
	if this.head == -1 {
		return false
	}
	if this.head == this.last {
		this.head = -1
		this.last = -1
		return true
	}

	if this.head == this.lastInd {
		this.head = 0
	} else {
		this.head += 1
	}
	return true
}


func (this *MyCircularDeque) DeleteLast() bool {
	if this.head == -1 && this.last == -1 {
		return false
	}
	if this.head == this.last {
		this.head = -1
		this.last = -1
		return true
	}

	if this.last == 0 {
		this.last = this.lastInd
	} else {
		this.last -= 1
	}
	return true
}


func (this *MyCircularDeque) GetFront() int {
	if this.head == -1 {
		return -1
	} else {
		return this.de[this.head]
	}
}


func (this *MyCircularDeque) GetRear() int {
	if this.last == -1 {
		return -1
	} else {
		return this.de[this.last]
	}
}


func (this *MyCircularDeque) IsEmpty() bool {
	if this.head == -1 {
		return true
	} else {
		return false
	}
}


func (this *MyCircularDeque) IsFull() bool {
	if this.head == -1 {
		return false
	}
	if this.last == this.lastInd {
		if this.head == 0 {
			return true
		} else {
			return false
		}
	} else {
		if this.last+1 == this.head {
			return true
		} else {
			return false
		}
	}
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */


