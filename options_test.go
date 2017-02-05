package wkhtmltopdf

import (
	"reflect"
	"testing"
)

func TestGlobalOptions(t *testing.T) {

	testcases := []struct {
		Options []Option
		Args    []string
	}{
		{[]Option{}, []string{}},
		{[]Option{NoCollate()}, []string{"--no-collate"}},
		{[]Option{DPI(300), Grayscale()}, []string{"--dpi", "300", "--grayscale"}},
		{[]Option{ImageDPI(500), ImageQuality(60)}, []string{"--image-dpi", "500", "--image-quality", "60"}},
		{[]Option{LowQuality(), CookieJar("testpath")}, []string{"--low-quality", "--cookie-jar", "testpath"}},
		{[]Option{MarginBottom("1cm"), MarginTop("2cm")}, []string{"--margin-bottom", "1cm", "--margin-top", "2cm"}},
		{[]Option{MarginLeft("15mm"), MarginRight("1.5cm")}, []string{"--margin-left", "15mm", "--margin-right", "1.5cm"}},
		{[]Option{Landscape(), PageHeight("20cm")}, []string{"--orientation", "landscape", "--page-height", "20cm"}},
		{[]Option{PageWidth("35cm")}, []string{"--page-width", "35cm"}},
		{[]Option{NoPDFCompression(), Quiet()}, []string{"--no-pdf-compression", "--quiet"}},
		{[]Option{Title("An excellent pdf")}, []string{"--title", "An excellent pdf"}},
		{[]Option{PageSize("A4")}, []string{"--page-size", "A4"}},
	}

	for _, tc := range testcases {
		doc := NewDocument(tc.Options...)
		if !reflect.DeepEqual(doc.options, tc.Args) {
			t.Errorf("Wrong arguments created. Expected: %v, Got: %v", tc.Args, doc.options)
		}
	}
}
