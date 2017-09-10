# Generalized Golang [Slice trick](https://github.com/golang/go/wiki/SliceTricks) with reflect
Are you ready to exchange runtime speed with ease of coding?

### Examples
```
import slice
...
  names := []string{"john", "ann", "diana", "sam", "anderson"}
  names = slice.Append(names, []string{"song", "choi"}).([]string)
  // or
  names = slice.AppendStr(names, []string{"song", "choi"})
```
```
type car struct {
  name string
  value int
}
...
	cars := []car{{"cadillac", 30}, {"jeep", 20}, {"sports", 100}}
	mycar_, cars_ := slice.Pop(cars)
	mycar := mycar_.(car)
	cars = cars_.([]car)
  // or
  mycar, cars := PopCar(cars)
  // but it needs regeneration with your custom struct definition in `slicetrick_generate.go` and `generate.go`
```
```
var In = slice.In
...
  numbers := []int {1, 35, 23, 29, 3, 7, 15}
  amILucky := In(numbers, 7)
```

### Result of benchmark as usage
 - Usually, switch-case style is a bit slower.
 - You can mock-up your program faster with reflection version of function.
 - When it's ready to ship your program, you can switch to type specified version of functions.
 - You can optimize further with custom list of precompiled types.

### How to change precompiled types list and regenerate
 - Open `generate.go` and edit type list (line number 2). Below is default value specifying all of the go value types.
```
//go:generate ./template -d bool int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr float32 float64 complex64 complex128 string
```
 - Type `go generate` to generate and build.

### Benchmark result

##### with all default types precompiled ( default setting of generate.go )

|                              | switch-case                       |                          | reflection |            |
|------------------------------|-----------------------------------|--------------------------|------------|------------|
| Function Name                | String Specified (ex:AppendStr()) | Reflection (ex:Append()) | String     | Reflection |
| Append                       | 244 ns/op                         | 625 ns/op                | 248 ns/op  | 600 ns/op  |
| Copy                         | 182 ns/op                         | 626 ns/op                | 174 ns/op  | 632 ns/op  |
| Cut                          | 134 ns/op                         | 532 ns/op                | 92.9 ns/op | 473 ns/op  |
| Delete                       | 105 ns/op                         | 512 ns/op                | 106 ns/op  | 477 ns/op  |
| DeleteWithoutPreservingOrder | 81.6 ns/op                        | 262 ns/op                | 80.3 ns/op | 288 ns/op  |
| Expand                       | 403 ns/op                         | 1184 ns/op               | 471 ns/op  | 1291 ns/op |
| Extend                       | 281 ns/op                         | 674 ns/op                | 286 ns/op  | 689 ns/op  |
| Insert                       | 295 ns/op                         | 1310 ns/op               | 297 ns/op  | 1252 ns/op |
| InsertVector                 | 303 ns/op                         | 1318 ns/op               | 326 ns/op  | 1321 ns/op |
| Pop                          | 84.2 ns/op                        | 307 ns/op                | 91.0 ns/op | 340 ns/op  |
| PopBack                      | 79.4 ns/op                        | 311 ns/op                | 92.8 ns/op | 314 ns/op  |
| Push                         | 185 ns/op                         | 487 ns/op                | 196 ns/op  | 500 ns/op  |
| PushBack                     | 241 ns/op                         | 642 ns/op                | 255 ns/op  | 656 ns/op  |
| Shift                        | 83.2 ns/op                        | 322 ns/op                | 83.4 ns/op | 316 ns/op  |
| Unshift                      | 264 ns/op                         | 673 ns/op                | 250 ns/op  | 707 ns/op  |
| FilterWithoutAlloc           | 289 ns/op                         | 2253 ns/op               | 275 ns/op  | 2118 ns/op |
| Reverse                      | 140 ns/op                         | 414 ns/op                | 130 ns/op  | 389 ns/op  |

##### with string only precompiled 

|                              | switch-case                       |                          | reflection |            |
|------------------------------|-----------------------------------|--------------------------|------------|------------|
| Function Name                | String Specified (ex:AppendStr()) | Reflection (ex:Append()) | String     | Reflection |
| Append                       | 246 ns/op                         | 597 ns/op                | 249 ns/op  | 620 ns/op  |
| Copy                         | 179 ns/op                         | 575 ns/op                | 198 ns/op  | 550 ns/op  |
| Cut                          | 109 ns/op                         | 540 ns/op                | 92.2 ns/op | 467 ns/op  |
| Delete                       | 104 ns/op                         | 524 ns/op                | 90.3 ns/op | 462 ns/op  |
| DeleteWithoutPreservingOrder | 87.9 ns/op                        | 268 ns/op                | 80.0 ns/op | 261 ns/op  |
| Expand                       | 416 ns/op                         | 1199 ns/op               | 396 ns/op  | 1186 ns/op |
| Extend                       | 323 ns/op                         | 715 ns/op                | 277 ns/op  | 673 ns/op  |
| Insert                       | 311 ns/op                         | 1310 ns/op               | 318 ns/op  | 1407 ns/op |
| InsertVector                 | 315 ns/op                         | 1273 ns/op               | 319 ns/op  | 1253 ns/op |
| Pop                          | 84.3 ns/op                        | 323 ns/op                | 85.8 ns/op | 298 ns/op  |
| PopBack                      | 83.3 ns/op                        | 320 ns/op                | 80.5 ns/op | 306 ns/op  |
| Push                         | 190 ns/op                         | 519 ns/op                | 184 ns/op  | 563 ns/op  |
| PushBack                     | 250 ns/op                         | 674 ns/op                | 264 ns/op  | 682 ns/op  |
| Shift                        | 85.2 ns/op                        | 319 ns/op                | 82.4 ns/op | 331 ns/op  |
| Unshift                      | 249 ns/op                         | 649 ns/op                | 248 ns/op  | 630 ns/op  |
| FilterWithoutAlloc           | 270 ns/op                         | 2038 ns/op               | 277 ns/op  | 2026 ns/op |
| Reverse                      | 138 ns/op                         | 427 ns/op                | 136 ns/op  | 380 ns/op  |

