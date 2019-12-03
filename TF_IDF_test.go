//Package tfidf Provides a simple implementation of Term-Frequency Inverse-Document-Frequency in golang

//Can return the points in tf values or tf-idf values

package tfidf

import (
	"reflect"
	"testing"

	kmeans "github.com/Nuno19/KMeans-Go"
)

func TestWordSet_Print(t *testing.T) {
	tests := []struct {
		name  string
		slice WordSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slice.Print()
		})
	}
}

func TestFloatMap_Print(t *testing.T) {
	tests := []struct {
		name string
		fmap FloatMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fmap.Print()
		})
	}
}

func TestWordCounts_Print(t *testing.T) {
	tests := []struct {
		name string
		fmap WordCounts
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fmap.Print()
		})
	}
}

func TestWordSet_Exists(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name  string
		slice WordSet
		args  args
		want  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slice.Exists(tt.args.word); got != tt.want {
				t.Errorf("WordSet.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordSet_ToLower(t *testing.T) {
	tests := []struct {
		name  string
		slice WordSet
		want  WordSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slice.ToLower(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordSet.ToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitCounts(t *testing.T) {
	type args struct {
		set WordSet
	}
	tests := []struct {
		name string
		args args
		want WordCounts
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitCounts(tt.args.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTF_IDF_SetCount(t *testing.T) {
	type args struct {
		corpus WordSet
	}
	tests := []struct {
		name   string
		tf_idf *TF_IDF
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tf_idf.SetCount(tt.args.corpus)
		})
	}
}

func TestTF_IDF_SetCountIdx(t *testing.T) {
	type args struct {
		corpus WordSet
		idx    int
	}
	tests := []struct {
		name   string
		tf_idf *TF_IDF
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tf_idf.SetCountIdx(tt.args.corpus, tt.args.idx)
		})
	}
}

func TestTF_IDF_GetSetCount(t *testing.T) {
	type args struct {
		corpus WordSet
	}
	tests := []struct {
		name   string
		tf_idf *TF_IDF
		args   args
		want   WordCounts
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tf_idf.GetSetCount(tt.args.corpus); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TF_IDF.GetSetCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTF_IDF_ComputeTF(t *testing.T) {
	type args struct {
		corpus WordSet
	}
	tests := []struct {
		name   string
		tf_idf *TF_IDF
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tf_idf.ComputeTF(tt.args.corpus)
		})
	}
}

func TestTF_IDF_GetComputedTF(t *testing.T) {
	type args struct {
		corpus WordSet
		counts WordCounts
	}
	tests := []struct {
		name   string
		tf_idf *TF_IDF
		args   args
		want   FloatMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tf_idf.GetComputedTF(tt.args.corpus, tt.args.counts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TF_IDF.GetComputedTF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTF_IDF_GetComputedIDF(t *testing.T) {
	type args struct {
		countMap WordCounts
	}
	tests := []struct {
		name   string
		tf_idf *TF_IDF
		args   args
		want   FloatMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tf_idf.GetComputedIDF(tt.args.countMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TF_IDF.GetComputedIDF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTF_IDF_GetComputedTFIDF(t *testing.T) {
	type args struct {
		tf  FloatMap
		idf FloatMap
	}
	tests := []struct {
		name   string
		tf_idf *TF_IDF
		args   args
		want   FloatMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tf_idf.GetComputedTFIDF(tt.args.tf, tt.args.idf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TF_IDF.GetComputedTFIDF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTF_IDF_ComputeIDF(t *testing.T) {
	tests := []struct {
		name   string
		tf_idf *TF_IDF
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tf_idf.ComputeIDF()
		})
	}
}

func TestTF_IDF_ComputeTFIDF(t *testing.T) {
	tests := []struct {
		name   string
		tf_idf *TF_IDF
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tf_idf.ComputeTFIDF()
		})
	}
}

func TestTF_IDF_AddToWordSet(t *testing.T) {
	type args struct {
		corpus []WordSet
	}
	tests := []struct {
		name   string
		tf_idf *TF_IDF
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tf_idf.AddToWordSet(tt.args.corpus)
		})
	}
}

func TestTF_IDF_SortMap(t *testing.T) {
	tests := []struct {
		name   string
		tf_idf *TF_IDF
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tf_idf.SortMap()
		})
	}
}

func TestTF_IDF_GetPointByIndexTF(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		name   string
		tf_idf *TF_IDF
		args   args
		want   kmeans.Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tf_idf.GetPointByIndexTF(tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TF_IDF.GetPointByIndexTF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTF_IDF_GetAllPointsTF(t *testing.T) {
	tests := []struct {
		name   string
		tf_idf *TF_IDF
		want   []kmeans.Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tf_idf.GetAllPointsTF(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TF_IDF.GetAllPointsTF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTF_IDF_GetPointByIndexTFIDF(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		name   string
		tf_idf *TF_IDF
		args   args
		want   kmeans.Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tf_idf.GetPointByIndexTFIDF(tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TF_IDF.GetPointByIndexTFIDF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTF_IDF_GetAllPointsTFIDF(t *testing.T) {
	tests := []struct {
		name   string
		tf_idf *TF_IDF
		want   []kmeans.Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tf_idf.GetAllPointsTFIDF(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TF_IDF.GetAllPointsTFIDF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTF_IDF_GetIDF(t *testing.T) {
	tests := []struct {
		name   string
		tf_idf *TF_IDF
		want   FloatMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tf_idf.GetIDF(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TF_IDF.GetIDF() = %v, want %v", got, tt.want)
			}
		})
	}
}
