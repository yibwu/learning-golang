package main

import (
	"errors"
	"fmt"
	"time"
)

type ReusableObj struct {

}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i += 1 {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObject(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <- time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObject(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func main()  {
	p := NewObjPool(10)

	for i := 0; i < 11; i += 1 {
		if obj, err := p.GetObject(time.Second); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(obj)
			p.ReleaseObject(obj)
		}
	}
}
