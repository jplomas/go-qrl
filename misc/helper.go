package misc

import (
	"github.com/theQRL/qrllib/goqrllib"
	"bytes"
	"container/list"
	"math"
)

type UcharVector struct {
	data goqrllib.UcharVector
}

func (v *UcharVector) AddBytes(data []byte) {
	for _, element := range data {
		v.data.Add(element)
	}
}

func (v *UcharVector) AddByte(data byte) {
	v.data.Add(data)
}

func (v *UcharVector) GetBytesBuffer() bytes.Buffer {
	var data bytes.Buffer
	for i := int64(0); i < v.data.Size(); i++ {
		value := v.data.Get(int(i))
		data.WriteByte(value)
	}
	return data
}

func (v *UcharVector) GetBytes() []byte {
	data := v.GetBytesBuffer()
	return data.Bytes()
}

func (v *UcharVector) GetString() string {
	data := v.GetBytesBuffer()
	return data.String()
}

func (v *UcharVector) GetData() goqrllib.UcharVector {
	return v.data
}

func (v *UcharVector) AddAt() goqrllib.UcharVector {
	return v.data
}

func (v *UcharVector) New(data goqrllib.UcharVector) {
	v.data = data
}

func BytesToUCharVector(data []byte) goqrllib.UcharVector {
	vector := goqrllib.NewUcharVector__SWIG_0()
	//z := *(*[]byte)(unsafe.Pointer(vector.Swigcptr()))
	//var sl = struct {
	//	addr uintptr
	//	len  int
	//	cap  int
	//}{ptr, length, length}
	for _, element := range data {
		vector.Add(element)
	}

	return vector
}

func UCharVectorToBytes(data goqrllib.UcharVector) []byte  {
	vector := UcharVector{}
	vector.New(data)

	return vector.GetBytes()
}

func UCharVectorToString(data goqrllib.UcharVector) string  {
	return string(UCharVectorToBytes(data))
}

func MerkleTXHash(hashes list.List) []byte {
	j := int(math.Ceil(math.Log2(float64(hashes.Len()))))
	var lArray list.List
	lArray.PushBack(hashes)
	for x := 0; x < j; x++ {
		var nextLayer list.List
		h := lArray.Back().Value.(list.List)
		i := h.Len() % 2 + h.Len() / 2
		e := h.Front()
		z := 0
		for k := 0; k < i; k++ {
			if h.Len() == z + 1 {
				nextLayer.PushBack(e.Value.([]byte))
			} else {
				tmp := UcharVector{}
				tmp.AddBytes(e.Value.([]byte))
				e := e.Next()
				tmp.AddBytes(e.Value.([]byte))
				nextLayer.PushBack(UCharVectorToBytes(goqrllib.Sha2_256(tmp.GetData())))
				e = e.Next()
			}
			z += 2
		}
		lArray.PushBack(nextLayer)
	}
	return lArray.Back().Value.(list.List).Back().Value.([]byte)
}

func Reverse(s [][]byte) [][]byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}