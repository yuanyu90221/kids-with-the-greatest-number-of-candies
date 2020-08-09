# kids-with-the-greatest-number-of-candies
## 題目解讀：

### 題目來源:

[kids-with-the-greatest-number-of-candies](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/)

### 原文:

Given the array candies and the integer extraCandies, where candies[i] represents the number of candies that the ith kid has.

For each kid check if there is a way to distribute extraCandies among the kids such that he or she can have the greatest number of candies among them. Notice that multiple kids can have the greatest number of candies.

### 解讀：

給一個正整數陣列candies 還有一個正整數 extraCandies

對於每個 candies[i] 代表 第i位 小孩所擁有的糖果

results 是一個布林值陣列

對於每個 results[i] 代表 第i位 小孩在給予多出來的 extraCandies

能不能成為其中最多糖果的人之一

這邊的最多可以同時存在好幾個

## 初步解法:
### 初步觀察:

首先是 如果原本就擁有最多糖果的小孩

在加上extraCandies一定是擁有最多的糖果

假設 目前擁有最多糖果的數量是x

每次檢查 candies[i] + extraCandies 是否大於或者等於 x

就知道是否能成為最多糖果擁有者之一

### 初步設計:

given an integer array candies and an integer extraCandies

create an integer array results with length of candies

set maximum = 0

loop index i = 0 to length of candies -1:

if candies[i] > maximum set maximum = candies[i]

loop index i = 0 to length of candies -1:

set results[i] = (candise[i] + extraCandies) > x 

return results
## 遇到的困難

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間

## 測資的撰寫

```golang
package kis_with_candies

import (
	"reflect"
	"testing"
)

func Test_kidsWithCandies(t *testing.T) {
	type args struct {
		candies      []int
		extraCandies int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "Example1",
			args: args{
				candies:      []int{2, 3, 5, 1, 3},
				extraCandies: 3,
			},
			want: []bool{true, true, true, false, true},
		},
		{
			name: "Example2",
			args: args{
				candies:      []int{4, 2, 1, 1, 2},
				extraCandies: 1,
			},
			want: []bool{true, false, false, false, false},
		},
		{
			name: "Example3",
			args: args{
				candies:      []int{12, 1, 12},
				extraCandies: 10,
			},
			want: []bool{true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kidsWithCandies(tt.args.candies, tt.args.extraCandies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kidsWithCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record
[leetcode ironman 30 day third -day](https://hackmd.io/yEl3bqA4SdqLChCb1woZaQ?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)