package sliceTrick_test

import (
	"fmt"
	slice "github.com/chidea/go-slice-trick"
	"strings"
	"testing"
)

var name string

var benchStr bool
var blankT *testing.T = &testing.T{}

func test(t *testing.T, testfn func(*testing.T), benchsfn, benchrfn func()) {
	if strings.HasPrefix(t.Name(), "Test") {
		testfn(t)
	} else if benchStr {
		benchsfn()
	} else {
		benchrfn()
	}
}
func test1(t *testing.T, tests func(a []string) []string, testr func(a interface{}) interface{}, arg []string) {
	test(t, func(t *testing.T) {
		logAndEval(t, tests(arg), testr(arg))
	}, func() {
		_ = tests(arg)
	}, func() {
		_ = testr(arg)
	})
}
func test1p(t *testing.T, tests func(a *[]string) []string, testr func(a interface{}) interface{}, arg []string) {
	test(t, func(t *testing.T) {
		a := slice.Copy(arg).([]string)
		b := slice.Copy(arg)
		logAndEval(t, tests(&a), testr(b))
	}, func() {
		//a := slice.Copy(arg).([]string)
		//_ = tests(&a)
		_ = tests(&arg)
	}, func() {
		//b := slice.Copy(arg)
		//_ = testr(b)
		_ = testr(&arg)
	})
}
func test1i(t *testing.T, tests func(a []string, i int) []string, testr func(a interface{}, i int) interface{}, a []string, i int) {
	test(t, func(t *testing.T) {
		logAndEval(t, tests(a, i), testr(a, i))
	}, func() {
		_ = tests(a, i)
	}, func() {
		_ = testr(a, i)
	})
}
func test1ij(t *testing.T, tests func(a []string, i, j int) []string, testr func(a interface{}, i, j int) interface{}, a []string, i, j int) {
	test(t, func(t *testing.T) {
		logAndEval(t, tests(a, i, j), testr(a, i, j))
	}, func() {
		_ = tests(a, i, j)
	}, func() {
		_ = testr(a, i, j)
	})
}
func test1v(t *testing.T, tests func(a []string, x string) []string, testr func(a, x interface{}) interface{}, a []string, x string) {
	test(t, func(t *testing.T) {
		logAndEval(t, tests(a, x), testr(a, x))
	}, func() {
		_ = tests(a, x)
	}, func() {
		_ = testr(a, x)
	})
}
func test1v_b(t *testing.T, tests func(a []string, x string) bool, testr func(a, x interface{}) bool, a []string, x string) {
	test(t, func(t *testing.T) {
		sval, rval := tests(a, x), testr(a, x)
		t.Log(sval)
		t.Log(rval)
		if sval != rval {
			t.Fail()
		}
	}, func() {
		_ = tests(a, x)
	}, func() {
		_ = testr(a, x)
	})
}
func test1vi(t *testing.T, tests func(a []string, x string, i int) []string, testr func(a, x interface{}, i int) interface{}, a []string, x string, i int) {
	test(t, func(t *testing.T) {
		logAndEval(t, tests(a, x, i), testr(a, x, i))
	}, func() {
		_ = tests(a, x, i)
	}, func() {
		_ = testr(a, x, i)
	})
}
func test1f(t *testing.T, tests func(a []string, f func(string) bool) []string, testr func(a, f interface{}) interface{}, a []string, fs func(string) bool) {
	test(t, func(t *testing.T) {
		logAndEval(t, tests(a, fs), testr(a, fs))
	}, func() {
		_ = tests(a, fs)
	}, func() {
		_ = testr(a, fs)
	})
}
func test12(t *testing.T, tests func(a []string) (string, []string), testr func(a interface{}) (interface{}, interface{}), a []string) {
	test(t, func(t *testing.T) {
		sa, sb := tests(a)
		ra, rb := testr(a)
		logAndEval(t, sa, ra)
		logAndEval(t, sb, rb)
	}, func() {
		_, _ = tests(a)
	}, func() {
		_, _ = testr(a)
	})
}
func test2(t *testing.T, tests func(a, b []string) []string, testr func(a, b interface{}) interface{}, a, b []string) {
	test(t, func(t *testing.T) {
		logAndEval(t, tests(a, b), testr(a, b))
	}, func() {
		_ = tests(a, b)
	}, func() {
		_ = testr(a, b)
	})
}
func test2i(t *testing.T, tests func(a, b []string, i int) []string, testr func(a, b interface{}, i int) interface{}, a, b []string, i int) {
	test(t, func(t *testing.T) {
		logAndEval(t, tests(a, b, i), testr(a, b, i))
	}, func() {
		_ = tests(a, b, i)
	}, func() {
		_ = testr(a, b, i)
	})
}

// Append
func TestAppend(t *testing.T) {
	test2(t, slice.AppendStr, slice.Append, []string{"a", "b", "c"}, []string{"d", "e", "f"})
}
func ExampleAppend() {
	fmt.Println(slice.Append([]string{"a", "b", "c"}, []string{"d", "e", "f"}))
	// Output : [a b c d e f]
}
func BenchmarkAppendStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestAppend(blankT)
	}
}
func BenchmarkAppendRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestAppend(blankT)
	}
}

// Copy
func TestCopy(t *testing.T) {
	test1(t, slice.CopyStr, slice.Copy, []string{"a", "b", "c"})
}

func ExampleCopy() {
	fmt.Println(slice.Copy([]string{"a", "b", "c"}))
	// Output : [a b c]
}
func BenchmarkCopyStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestCopy(blankT)
	}
}

func BenchmarkCopyRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestCopy(blankT)
	}
}

// Cut
func TestCut(t *testing.T) {
	test1ij(t, slice.CutStr, slice.Cut, []string{"a", "b", "c"}, 1, 2)
}

func ExampleCut() {
	fmt.Println(slice.Cut([]string{"a", "b", "c"}, 1, 3))
	// Output : [a]
}

func BenchmarkCutStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestCut(blankT)
	}
}

func BenchmarkCutRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestCut(blankT)
	}
}

// Delete
func TestDelete(t *testing.T) {
	test1i(t, slice.DeleteStr, slice.Delete, []string{"a", "b", "c"}, 1)
}

func ExampleDelete() {
	fmt.Println(slice.Delete([]string{"a", "b", "c"}, 1))
	// Output : [a c]
}

func BenchmarkDeleteStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestDelete(blankT)
	}
}

func BenchmarkDeleteRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestDelete(blankT)
	}
}

// DeleteWithoutPreservingOrder
func TestDeleteWithoutPreservingOrder(t *testing.T) {
	test1i(t, slice.DeleteWithoutPreservingOrderStr, slice.DeleteWithoutPreservingOrder, []string{"a", "b", "c"}, 1)
}

func ExampleDeleteWithoutPreservingOrder() {
	fmt.Println(slice.DeleteWithoutPreservingOrder([]string{"a", "b", "c"}, 1))
	// Output : [a c]
}

func BenchmarkDeleteWithoutPreservingOrderStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestDeleteWithoutPreservingOrder(blankT)
	}
}

func BenchmarkDeleteWithoutPreservingOrderRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestDeleteWithoutPreservingOrder(blankT)
	}
}

// Expand
func TestExpand(t *testing.T) {
	test1ij(t, slice.ExpandStr, slice.Expand, []string{"a", "b", "c"}, 1, 3)
}
func ExampleExpand() {
	fmt.Println(slice.Expand([]string{"a", "b", "c"}, 1, 3))
	// Output : [a    b c]
}

func BenchmarkExpandStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestExpand(blankT)
	}
}

func BenchmarkExpandRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestExpand(blankT)
	}
}

// Extend
func TestExtend(t *testing.T) {
	test1i(t, slice.ExtendStr, slice.Extend, []string{"a", "b", "c"}, 3)
}
func ExampleExtend() {
	fmt.Println(slice.Extend([]string{"a", "b", "c"}, 3))
	// Output : [a b c    ]
}

func BenchmarkExtendStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestExtend(blankT)
	}
}

func BenchmarkExtendRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestExtend(blankT)
	}
}

// Insert
func TestInsert(t *testing.T) {
	test1vi(t, slice.InsertStr, slice.Insert, []string{"a", "c"}, "b", 1)
}
func ExampleInsert() {
	fmt.Println(slice.Expand([]string{"a", "b", "c"}, 1, 3))
	// Output : [a b c]
}

func BenchmarkInsertStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestInsert(blankT)
	}
}

func BenchmarkInsertRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestInsert(blankT)
	}
}

// InsertVector
func TestInsertVector(t *testing.T) {
	test2i(t, slice.InsertVectorStr, slice.InsertVector, []string{"a", "d"}, []string{"b", "c"}, 1)
}
func ExampleInsertVector() {
	fmt.Println(slice.InsertVector([]string{"a", "d"}, []string{"b", "c"}, 1))
	// Output : [a b c d]
}

func BenchmarkInsertVectorStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestInsert(blankT)
	}
}
func BenchmarkInsertVectorRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestInsert(blankT)
	}
}

func TestPop(t *testing.T) {
	test12(t, slice.PopStr, slice.Pop, []string{"a", "b", "c"})
}
func ExamplePop() {
	fmt.Println(slice.Pop([]string{"a", "b", "c"}))
	// Output :
	// c
	// [a b]
}

func BenchmarkPopStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestPop(blankT)
	}
}
func BenchmarkPopRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestPop(blankT)
	}
}

func TestPopBack(t *testing.T) {
	test12(t, slice.PopBackStr, slice.PopBack, []string{"a", "b", "c"})
}
func ExamplePopBack() {
	fmt.Println(slice.PopBack([]string{"a", "b", "c"}))
	// Output :
	// a
	// [b c]
}

func BenchmarkPopBackStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestPopBack(blankT)
	}
}
func BenchmarkPopBackRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestPopBack(blankT)
	}
}
func TestPush(t *testing.T) {
	test1v(t, slice.PushStr, slice.Push, []string{"a", "b"}, "c")
}
func ExamplePush() {
	fmt.Println(slice.Push([]string{"a", "b"}, "c"))
	// Output : [a b c]
}

func BenchmarkPushStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestPush(blankT)
	}
}
func BenchmarkPushRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestPush(blankT)
	}
}
func TestPushBack(t *testing.T) {
	test1v(t, slice.PushBackStr, slice.PushBack, []string{"b", "c"}, "a")
}
func ExamplePushBack() {
	fmt.Println(slice.PushBack([]string{"b", "c"}, "a"))
	// Output : [a b c]
}

func BenchmarkPushBackStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestPushBack(blankT)
	}
}
func BenchmarkPushBackRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestPushBack(blankT)
	}
}
func TestShift(t *testing.T) {
	test12(t, slice.ShiftStr, slice.Shift, []string{"a", "b", "c"})
}
func ExampleShift() {
	fmt.Println(slice.PopBack([]string{"a", "b", "c"}))
	// Output :
	// a
	// [b c]
}

func BenchmarkShiftStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestShift(blankT)
	}
}
func BenchmarkShiftRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestShift(blankT)
	}
}
func TestUnshift(t *testing.T) {
	test1v(t, slice.UnshiftStr, slice.Unshift, []string{"b", "c"}, "a")
}
func ExampleUnshift() {
	fmt.Println(slice.Unshift([]string{"b", "c"}, "a"))
	// Output : [a b c]
}

func BenchmarkUnshiftStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestUnshift(blankT)
	}
}
func BenchmarkUnshiftRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestUnshift(blankT)
	}
}
func TestFilterWithoutAlloc(t *testing.T) {
	test1f(t, slice.FilterWithoutAllocStr, slice.FilterWithoutAlloc, []string{"a", "b", "c"}, func(o string) bool {
		return o == "b" || o == "c"
	})
}
func ExampleFilterWithoutAlloc() {
	fmt.Println(slice.FilterWithoutAlloc([]string{"a", "b", "c"},
		func(o interface{}) bool {
			v := o.(string)
			return v == "b" || v == "c"
		}))
	// Output : [b c]
}

func BenchmarkFilterWithoutAllocStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestFilterWithoutAlloc(blankT)
	}
}
func BenchmarkFilterWithoutAllocRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestFilterWithoutAlloc(blankT)
	}
}
func TestReverse(t *testing.T) {
	test1p(t, slice.ReverseStr, slice.Reverse, []string{"a", "b", "c"})
}
func ExampleReverse() {
	fmt.Println(slice.Reverse([]string{"a", "b", "c"}))
	// Output : [c b a]
}

func BenchmarkReverseStr(b *testing.B) {
	benchStr = true
	for i := 0; i < b.N; i++ {
		TestReverse(blankT)
	}
}
func BenchmarkReverseRef(b *testing.B) {
	benchStr = false
	for i := 0; i < b.N; i++ {
		TestReverse(blankT)
	}
}

func TestIn(t *testing.T) {
	test1v_b(t, slice.InStr, slice.In, []string{"a", "b", "c"}, "a")
}
func TestMethodIn(t *testing.T) {
	v := slice.T{"b"}.In([]string{"a", "b", "c"})
	t.Log(v)
	if !v {
		t.Fail()
	}
}
func logAndEval(t *testing.T, ts, tr interface{}) {
	t.Logf("%+v", ts)
	t.Logf("%+v", tr)
	eval(t, ts, tr)

	/*func logAndEval(tb interface{}, rst ...interface{}) {
	switch tb.(type) {
	case *testing.T:
		for i := 0; i < len(rst); i++ {
			log.Printf("%v", rst[i])
		}
		eval(tb.(*testing.T), rst[0], rst[1])
	}*/
}

func eval(t *testing.T, ts, tr interface{}) {
	if !slice.Equal(ts, tr) {
		if l := slice.Len(ts); l != slice.Len(tr) {
			t.Logf("length of testString is %d and length of testReflect %d", l, slice.Len(tr))
			t.Fail()
		} else {
			for i := 0; i < l; i++ {
				if ts.([]string)[i] != tr.([]string)[i] {
					t.Logf("index %d of testString is %s but testReflect is %s", i, ts.([]string)[i], tr.([]string)[i])
					t.Fail()
				}
			}
		}
		t.Fail()
	}
}

func TestNames(t *testing.T) {
	names := []string{"john", "ann", "diana", "sam", "anderson"}
	names = slice.Append(names, []string{"song", "choi"}).([]string)
	t.Log(names)
	// or
	names = slice.AppendStr(names, []string{"song", "choi"})
	t.Log(names)
}

type car struct {
	name  string
	value int
}

func TestCar(t *testing.T) {
	cars := []car{{"cadillac", 30}, {"jeep", 20}, {"sports", 100}}
	mycar_, cars_ := slice.Pop(cars)
	mycar := mycar_.(car)
	cars = cars_.([]car)
	t.Log(mycar)
	// or
	//mycar, cars := PopCar(cars)
	// but it needs regeneration with your custom struct definition in `slicetrick_generate.go` and `generate.go`
}

var In = slice.In

func TestLuck(t *testing.T) {
	numbers := []int{1, 35, 23, 29, 3, 7, 15}
	amILucky := In(numbers, 7)
	t.Log(amILucky)
}
