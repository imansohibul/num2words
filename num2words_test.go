package num2words_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"num2words"
)

var tests = []struct {
	Number int64
	Words  string
}{
	{0, "nol"},
	{1, "satu"},
	{2, "dua"},
	{3, "tiga"},
	{4, "empat"},
	{5, "lima"},
	{6, "enam"},
	{7, "tujuh"},
	{8, "delapan"},
	{9, "sembilan"},
	{10, "sepuluh"},
	{11, "sebelas"},
	{12, "dua belas"},
	{13, "tiga belas"},
	{14, "empat belas"},
	{15, "lima belas"},
	{16, "enam belas"},
	{17, "tujuh belas"},
	{18, "delapan belas"},
	{19, "sembilan belas"},
	{20, "dua puluh"},
	{21, "dua puluh satu"},
	{25, "dua puluh lima"},
	{29, "dua puluh sembilan"},
	{30, "tiga puluh"},
	{32, "tiga puluh dua"},
	{36, "tiga puluh enam"},
	{40, "empat puluh"},
	{43, "empat puluh tiga"},
	{44, "empat puluh empat"},
	{50, "lima puluh"},
	{60, "enam puluh"},
	{67, "enam puluh tujuh"},
	{70, "tujuh puluh"},
	{78, "tujuh puluh delapan"},
	{80, "delapan puluh"},
	{90, "sembilan puluh"},
	{100, "seratus"},
	{101, "seratus satu"},
	{105, "seratus lima"},
	{110, "seratus sepuluh"},
	{120, "seratus dua puluh"},
	{189, "seratus delapan puluh sembilan"},
	{200, "dua ratus"},
	{267, "dua ratus enam puluh tujuh"},
	{999, "sembilan ratus sembilan puluh sembilan"},
	{1000, "seribu"},
	{1001, "seribu satu"},
	{1005, "seribu lima"},
	{1010, "seribu sepuluh"},
	{1110, "seribu seratus sepuluh"},
	{2902, "dua ribu sembilan ratus dua"},
	{9999, "sembilan ribu sembilan ratus sembilan puluh sembilan"},
	{100000, "seratus ribu"},
	{100001, "seratus ribu satu"},
	{100010, "seratus ribu sepuluh"},
	{100019, "seratus ribu sembilan belas"},
	{578494, "lima ratus tujuh puluh delapan ribu empat ratus sembilan puluh empat"},
	{999999, "sembilan ratus sembilan puluh sembilan ribu sembilan ratus sembilan puluh sembilan"},
	{1000001, "satu juta satu"},
	{1000010, "satu juta sepuluh"},
	{2000120, "dua juta seratus dua puluh"},
	{2550000, "dua juta lima ratus lima puluh ribu"},
	{5892000, "lima juta delapan ratus sembilan puluh dua ribu"},
	{9999999, "sembilan juta sembilan ratus sembilan puluh sembilan ribu sembilan ratus sembilan puluh sembilan"},
	{10000000, "sepuluh juta"},
	{10000001, "sepuluh juta satu"},
	{10000010, "sepuluh juta sepuluh"},
	{10000120, "sepuluh juta seratus dua puluh"},
	{24900000, "dua puluh empat juta sembilan ratus ribu"},
	{99999999, "sembilan puluh sembilan juta sembilan ratus sembilan puluh sembilan ribu sembilan ratus sembilan puluh sembilan"},
	{100000000, "seratus juta"},
	{100000001, "seratus juta satu"},
	{879560000, "delapan ratus tujuh puluh sembilan juta lima ratus enam puluh ribu"},
	{879569898, "delapan ratus tujuh puluh sembilan juta lima ratus enam puluh sembilan ribu delapan ratus sembilan puluh delapan"},
	{2000000000, "dua miliar"},
	{2000000001, "dua miliar satu"},
	{23000000100, "dua puluh tiga miliar seratus"},
	{892000000000, "delapan ratus sembilan puluh dua miliar"},
	{999940000000, "sembilan ratus sembilan puluh sembilan miliar sembilan ratus empat puluh juta"},
	{999949999999, "sembilan ratus sembilan puluh sembilan miliar sembilan ratus empat puluh sembilan juta sembilan ratus sembilan puluh sembilan ribu sembilan ratus sembilan puluh sembilan"},
	{1000000000000, "satu triliun"},
	{970000000000000, "sembilan ratus tujuh puluh triliun"},
	{970237000000000, "sembilan ratus tujuh puluh triliun dua ratus tiga puluh tujuh miliar"},
	{970237255000000, "sembilan ratus tujuh puluh triliun dua ratus tiga puluh tujuh miliar dua ratus lima puluh lima juta"},
	{970237255258000, "sembilan ratus tujuh puluh triliun dua ratus tiga puluh tujuh miliar dua ratus lima puluh lima juta dua ratus lima puluh delapan ribu"},
	{970237255295800, "sembilan ratus tujuh puluh triliun dua ratus tiga puluh tujuh miliar dua ratus lima puluh lima juta dua ratus sembilan puluh lima ribu delapan ratus"},
	{970237255295823, "sembilan ratus tujuh puluh triliun dua ratus tiga puluh tujuh miliar dua ratus lima puluh lima juta dua ratus sembilan puluh lima ribu delapan ratus dua puluh tiga"},
}

func TestConvert(t *testing.T) {
	t.Parallel()
	for _, tt := range tests {
		t.Run("testconvert", func(t *testing.T) {
			w, err := num2words.Convert(tt.Number)
			assert.Nil(t, err)
			assert.Equal(t, tt.Words, w)
		})
	}

}

func TestConvertMinus(t *testing.T) {
	t.Parallel()
	for _, tt := range tests {
		t.Run("testconvertminus", func(t *testing.T) {
			w, err := num2words.Convert(tt.Number * -1)
			assert.Nil(t, err)
			if tt.Number == 0 {
				assert.Equal(t, tt.Words, w)
			} else {
				assert.Equal(t, "negatif "+tt.Words, w)
			}

		})
	}
}

func TestConvertError(t *testing.T) {
	w, err := num2words.Convert(1000000000000000)
	assert.NotNil(t, err)
	assert.Equal(t, "", w)
}
