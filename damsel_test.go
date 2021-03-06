package damsel

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"dasa.cc/damsel/parse"
)

var TestsDir = "tests"

func execute(t *Template, data interface{}) (string, error) {
	tpl := NewHtmlTemplate(t)
	r, err := tpl.Execute(data)
	if err != nil {
		return "", err
	}
	return r, nil
}

func get_html(t *testing.T, s string) string {
	b, err := ioutil.ReadFile(filepath.Join(TestsDir, s+".html"))
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(string(b))
}

func test(t *testing.T, s string, data interface{}) {
	fmt.Println("testing:", s)
	TemplateDir = TestsDir
	html := get_html(t, s)

	tpl, _ := ParseFile(s + ".dmsl")
	r, err := execute(tpl, data)
	r = strings.TrimSpace(r)
	if r != html {
		fmt.Println("\nExpected\n========\n", html, "\nReceived\n========\n", r, "\n\n")
		t.Fatal("parse failed:", s, "\n", err)
	}
}

func Test_html(t *testing.T) {
	test(t, "html", nil)
}

func Test_indent(t *testing.T) {
	test(t, "indent", nil)
}

func Test_variable_indent(t *testing.T) {
	test(t, "variable_indent", nil)
}

func Test_inline(t *testing.T) {
	data := []string{"a", "b", "c", "d"}
	test(t, "inline", data)
}

func Test_multiline_text(t *testing.T) {
	test(t, "multiline_text", nil)
}

func Test_ul_range_header(t *testing.T) {
	data := []string{"a", "b", "c", "d"}
	test(t, "ul_range_header", data)
}

func Test_tabs(t *testing.T) {
	test(t, "tabs", nil)
}

func Test_tag_hashes(t *testing.T) {
	test(t, "tag_hashes", nil)
}

func Test_html_comment(t *testing.T) {
	test(t, "hcomment", nil)
}

func Test_extends(t *testing.T) {
	test(t, "extends", nil)
}

func Test_extends_super(t *testing.T) {
	test(t, "extends_super", nil)
}

func Test_big_table(t *testing.T) {
	table := [2][10]int{}
	test(t, "bigtable", table)
}

func Benchmark_parser(b *testing.B) {
	b.StopTimer()
	bytes, err := ioutil.ReadFile(filepath.Join(TestsDir, "bigtable2.dmsl"))
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		parse.DocParse(bytes)
	}
}

func Benchmark_bigtable_go_stdlib(b *testing.B) {
	b.StopTimer()
	bytes, err := ioutil.ReadFile(filepath.Join(TestsDir, "bigtable.dmsl"))
	if err != nil {
		b.Fatal(err)
	}
	table := [1000][10]int{}
	tmpl, _ := Parse(bytes)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		execute(tmpl, table)
	}
}

func Benchmark_bigtable(b *testing.B) {
	b.StopTimer()
	bytes, err := ioutil.ReadFile(filepath.Join(TestsDir, "bigtable.dmsl"))
	if err != nil {
		b.Fatal(err)
	}
	table := [1000][10]int{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tpl, _ := Parse(bytes)
		execute(tpl, table)
	}
}

func Benchmark_bigtable2(b *testing.B) {
	b.StopTimer()
	bytes, err := ioutil.ReadFile(filepath.Join(TestsDir, "bigtable2.dmsl"))
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tpl, _ := Parse(bytes)
		execute(tpl, nil)
	}
}
