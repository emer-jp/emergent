// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package params

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/goki/gi/gi"
	"github.com/goki/ki/indent"
	"github.com/goki/ki/ki"
)

// WriteGoPrelude writes the start of a go file in package main that starts a
// variable assignment to given variable -- for start of SaveGoCode methods.
func WriteGoPrelude(w io.Writer, varNm string) {
	w.Write([]byte("// File generated by params.SaveGoCode\n\n"))
	w.Write([]byte("package main\n\n"))
	w.Write([]byte(`import "github.com/emer/emergent/params"`))
	w.Write([]byte("\n\nvar " + varNm + " = "))
}

// OpenJSON opens params from a JSON-formatted file.
func (pr *Params) OpenJSON(filename gi.FileName) error {
	*pr = make(Params) // reset
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "File Not Found", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

// SaveJSON saves params to a JSON-formatted file.
func (pr *Params) SaveJSON(filename gi.FileName) error {
	b, err := json.MarshalIndent(pr, "", "  ")
	if err != nil {
		log.Println(err) // unlikely
		return err
	}
	err = ioutil.WriteFile(string(filename), b, 0644)
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
	}
	return err
}

// WriteGoCode writes params to corresponding Go initializer code.
func (pr *Params) WriteGoCode(w io.Writer, depth int) {
	w.Write([]byte(fmt.Sprintf("params.Params{\n")))
	depth++
	paths := make([]string, len(*pr)) // alpha-sort paths for consistent output
	ctr := 0
	for pt := range *pr {
		paths[ctr] = pt
		ctr++
	}
	sort.StringSlice(paths).Sort()
	for _, pt := range paths {
		pv := (*pr)[pt]
		w.Write(indent.TabBytes(depth))
		w.Write([]byte(fmt.Sprintf("%q: %q,\n", pt, pv)))
	}
	depth--
	w.Write(indent.TabBytes(depth))
	w.Write([]byte("}"))
}

// StringGoCode returns Go initializer code as a byte string.
func (pr *Params) StringGoCode() []byte {
	var buf bytes.Buffer
	pr.WriteGoCode(&buf, 0)
	return buf.Bytes()
}

// SaveGoCode saves params to corresponding Go initializer code.
func (pr *Params) SaveGoCode(filename gi.FileName) error {
	fp, err := os.Create(string(filename))
	defer fp.Close()
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	WriteGoPrelude(fp, "SavedParams")
	pr.WriteGoCode(fp, 0)
	return nil
}

/////////////////////////////////////////////////////////
//   Sel

// OpenJSON opens params from a JSON-formatted file.
func (pr *Sel) OpenJSON(filename gi.FileName) error {
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "File Not Found", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

// SaveJSON saves params to a JSON-formatted file.
func (pr *Sel) SaveJSON(filename gi.FileName) error {
	b, err := json.MarshalIndent(pr, "", "  ")
	if err != nil {
		log.Println(err) // unlikely
		return err
	}
	err = ioutil.WriteFile(string(filename), b, 0644)
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
	}
	return err
}

// WriteGoCode writes params to corresponding Go initializer code.
func (pr *Sel) WriteGoCode(w io.Writer, depth int) {
	w.Write([]byte(fmt.Sprintf("Sel: %q, Desc: %q,\n", pr.Sel, pr.Desc)))
	depth++
	w.Write(indent.TabBytes(depth))
	w.Write([]byte("Params: "))
	pr.Params.WriteGoCode(w, depth)
}

// StringGoCode returns Go initializer code as a byte string.
func (pr *Sel) StringGoCode() []byte {
	var buf bytes.Buffer
	pr.WriteGoCode(&buf, 0)
	return buf.Bytes()
}

// SaveGoCode saves params to corresponding Go initializer code.
func (pr *Sel) SaveGoCode(filename gi.FileName) error {
	fp, err := os.Create(string(filename))
	defer fp.Close()
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	WriteGoPrelude(fp, "SavedParamsSel")
	pr.WriteGoCode(fp, 0)
	return nil
}

/////////////////////////////////////////////////////////
//   Sheet

// OpenJSON opens params from a JSON-formatted file.
func (pr *Sheet) OpenJSON(filename gi.FileName) error {
	*pr = make(Sheet, 0) // reset
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "File Not Found", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

// SaveJSON saves params to a JSON-formatted file.
func (pr *Sheet) SaveJSON(filename gi.FileName) error {
	b, err := json.MarshalIndent(pr, "", "  ")
	if err != nil {
		log.Println(err) // unlikely
		return err
	}
	err = ioutil.WriteFile(string(filename), b, 0644)
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
	}
	return err
}

// WriteGoCode writes params to corresponding Go initializer code.
func (pr *Sheet) WriteGoCode(w io.Writer, depth int) {
	w.Write([]byte(fmt.Sprintf("params.Sheet{\n")))
	depth++
	for _, pv := range *pr {
		w.Write(indent.TabBytes(depth))
		w.Write([]byte("{"))
		pv.WriteGoCode(w, depth)
		w.Write([]byte("},\n"))
	}
	depth--
	w.Write(indent.TabBytes(depth))
	w.Write([]byte("},\n"))
}

// StringGoCode returns Go initializer code as a byte string.
func (pr *Sheet) StringGoCode() []byte {
	var buf bytes.Buffer
	pr.WriteGoCode(&buf, 0)
	return buf.Bytes()
}

// SaveGoCode saves params to corresponding Go initializer code.
func (pr *Sheet) SaveGoCode(filename gi.FileName) error {
	fp, err := os.Create(string(filename))
	defer fp.Close()
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	WriteGoPrelude(fp, "SavedParamsSheet")
	pr.WriteGoCode(fp, 0)
	return nil
}

/////////////////////////////////////////////////////////
//   Sheets

// OpenJSON opens params from a JSON-formatted file.
func (pr *Sheets) OpenJSON(filename gi.FileName) error {
	*pr = make(Sheets) // reset
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "File Not Found", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

// SaveJSON saves params to a JSON-formatted file.
func (pr *Sheets) SaveJSON(filename gi.FileName) error {
	b, err := json.MarshalIndent(pr, "", "  ")
	if err != nil {
		log.Println(err) // unlikely
		return err
	}
	err = ioutil.WriteFile(string(filename), b, 0644)
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
	}
	return err
}

// WriteGoCode writes params to corresponding Go initializer code.
func (pr *Sheets) WriteGoCode(w io.Writer, depth int) {
	w.Write([]byte(fmt.Sprintf("params.Sheets{\n")))
	depth++
	nms := make([]string, len(*pr)) // alpha-sort names for consistent output
	ctr := 0
	for nm := range *pr {
		nms[ctr] = nm
		ctr++
	}
	sort.StringSlice(nms).Sort()
	for _, nm := range nms {
		pv := (*pr)[nm]
		w.Write(indent.TabBytes(depth))
		w.Write([]byte(fmt.Sprintf("%q: &", nm)))
		pv.WriteGoCode(w, depth)
	}
	depth--
	w.Write(indent.TabBytes(depth))
	w.Write([]byte("}"))
}

// StringGoCode returns Go initializer code as a byte string.
func (pr *Sheets) StringGoCode() []byte {
	var buf bytes.Buffer
	pr.WriteGoCode(&buf, 0)
	return buf.Bytes()
}

// SaveGoCode saves params to corresponding Go initializer code.
func (pr *Sheets) SaveGoCode(filename gi.FileName) error {
	fp, err := os.Create(string(filename))
	defer fp.Close()
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	WriteGoPrelude(fp, "SavedParamsSheets")
	pr.WriteGoCode(fp, 0)
	return nil
}

/////////////////////////////////////////////////////////
//   Set

// OpenJSON opens params from a JSON-formatted file.
func (pr *Set) OpenJSON(filename gi.FileName) error {
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "File Not Found", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

// SaveJSON saves params to a JSON-formatted file.
func (pr *Set) SaveJSON(filename gi.FileName) error {
	b, err := json.MarshalIndent(pr, "", "  ")
	if err != nil {
		log.Println(err) // unlikely
		return err
	}
	err = ioutil.WriteFile(string(filename), b, 0644)
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
	}
	return err
}

// WriteGoCode writes params to corresponding Go initializer code.
func (pr *Set) WriteGoCode(w io.Writer, depth int) {
	w.Write([]byte(fmt.Sprintf("Name: %q, Desc: %q, Sheets: ", pr.Name, pr.Desc)))
	pr.Sheets.WriteGoCode(w, depth)
}

// StringGoCode returns Go initializer code as a byte string.
func (pr *Set) StringGoCode() []byte {
	var buf bytes.Buffer
	pr.WriteGoCode(&buf, 0)
	return buf.Bytes()
}

// SaveGoCode saves params to corresponding Go initializer code.
func (pr *Set) SaveGoCode(filename gi.FileName) error {
	fp, err := os.Create(string(filename))
	defer fp.Close()
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	WriteGoPrelude(fp, "SavedParamsSet")
	pr.WriteGoCode(fp, 0)
	return nil
}

/////////////////////////////////////////////////////////
//   Sets

// OpenJSON opens params from a JSON-formatted file.
func (pr *Sets) OpenJSON(filename gi.FileName) error {
	*pr = make(Sets, 0, 10) // reset
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "File Not Found", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

// SaveJSON saves params to a JSON-formatted file.
func (pr *Sets) SaveJSON(filename gi.FileName) error {
	b, err := json.MarshalIndent(pr, "", "  ")
	if err != nil {
		log.Println(err) // unlikely
		return err
	}
	err = ioutil.WriteFile(string(filename), b, 0644)
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
	}
	return err
}

// WriteGoCode writes params to corresponding Go initializer code.
func (pr *Sets) WriteGoCode(w io.Writer, depth int) {
	w.Write([]byte(fmt.Sprintf("params.Sets{\n")))
	depth++
	for _, st := range *pr {
		w.Write(indent.TabBytes(depth))
		w.Write([]byte("{"))
		st.WriteGoCode(w, depth)
		w.Write([]byte("},\n"))
	}
	depth--
	w.Write(indent.TabBytes(depth))
	w.Write([]byte("}\n"))
}

// StringGoCode returns Go initializer code as a byte string.
func (pr *Sets) StringGoCode() []byte {
	var buf bytes.Buffer
	pr.WriteGoCode(&buf, 0)
	return buf.Bytes()
}

// SaveGoCode saves params to corresponding Go initializer code.
func (pr *Sets) SaveGoCode(filename gi.FileName) error {
	fp, err := os.Create(string(filename))
	defer fp.Close()
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	WriteGoPrelude(fp, "SavedParamsSets")
	pr.WriteGoCode(fp, 0)
	return nil
}

var ParamsProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"SaveJSON", ki.Props{
			"label": "Save As...",
			"desc":  "save to JSON formatted file",
			"icon":  "file-save",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".params",
				}},
			},
		}},
		{"OpenJSON", ki.Props{
			"label": "Open...",
			"desc":  "open from JSON formatted file",
			"icon":  "file-open",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".params",
				}},
			},
		}},
		{"sep-gocode", ki.BlankProp{}},
		{"SaveGoCode", ki.Props{
			"label": "Save Code As...",
			"desc":  "save to Go-formatted initializer code in file",
			"icon":  "go",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".go",
				}},
			},
		}},
		{"StringGoCode", ki.Props{
			"label":       "Show Code",
			"desc":        "shows the Go-formatted initializer code, can be copy / pasted into program",
			"icon":        "go",
			"show-return": true,
		}},
	},
}

var SelProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"SaveJSON", ki.Props{
			"label": "Save As...",
			"desc":  "save to JSON formatted file",
			"icon":  "file-save",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".params",
				}},
			},
		}},
		{"OpenJSON", ki.Props{
			"label": "Open...",
			"desc":  "open from JSON formatted file",
			"icon":  "file-open",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".params",
				}},
			},
		}},
		{"sep-gocode", ki.BlankProp{}},
		{"SaveGoCode", ki.Props{
			"label": "Save Code As...",
			"desc":  "save to Go-formatted initializer code in file",
			"icon":  "go",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".go",
				}},
			},
		}},
		{"StringGoCode", ki.Props{
			"label":       "Show Code",
			"desc":        "shows the Go-formatted initializer code, can be copy / pasted into program",
			"icon":        "go",
			"show-return": true,
		}},
	},
}

var SheetProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"SaveJSON", ki.Props{
			"label": "Save As...",
			"desc":  "save to JSON formatted file",
			"icon":  "file-save",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".params",
				}},
			},
		}},
		{"OpenJSON", ki.Props{
			"label": "Open...",
			"desc":  "open from JSON formatted file",
			"icon":  "file-open",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".params",
				}},
			},
		}},
		{"sep-gocode", ki.BlankProp{}},
		{"SaveGoCode", ki.Props{
			"label": "Save Code As...",
			"desc":  "save to Go-formatted initializer code in file",
			"icon":  "go",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".go",
				}},
			},
		}},
		{"StringGoCode", ki.Props{
			"label":       "Show Code",
			"desc":        "shows the Go-formatted initializer code, can be copy / pasted into program",
			"icon":        "go",
			"show-return": true,
		}},
	},
}

var SheetsProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"SaveJSON", ki.Props{
			"label": "Save As...",
			"desc":  "save to JSON formatted file",
			"icon":  "file-save",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".params",
				}},
			},
		}},
		{"OpenJSON", ki.Props{
			"label": "Open...",
			"desc":  "open from JSON formatted file",
			"icon":  "file-open",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".params",
				}},
			},
		}},
		{"sep-gocode", ki.BlankProp{}},
		{"SaveGoCode", ki.Props{
			"label": "Save Code As...",
			"desc":  "save to Go-formatted initializer code in file",
			"icon":  "go",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".go",
				}},
			},
		}},
		{"StringGoCode", ki.Props{
			"label":       "Show Code",
			"desc":        "shows the Go-formatted initializer code, can be copy / pasted into program",
			"icon":        "go",
			"show-return": true,
		}},
		{"sep-diffs", ki.BlankProp{}},
		{"DiffsWithin", ki.Props{
			"desc":        "reports where the same param path is being set to different values within this set (both within the same Sheet and betwen sheets)",
			"icon":        "search",
			"show-return": true,
		}},
	},
}

var SetProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"SaveJSON", ki.Props{
			"label": "Save As...",
			"desc":  "save to JSON formatted file",
			"icon":  "file-save",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".params",
				}},
			},
		}},
		{"OpenJSON", ki.Props{
			"label": "Open...",
			"desc":  "open from JSON formatted file",
			"icon":  "file-open",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".params",
				}},
			},
		}},
		{"sep-gocode", ki.BlankProp{}},
		{"SaveGoCode", ki.Props{
			"label": "Save Code As...",
			"desc":  "save to Go-formatted initializer code in file",
			"icon":  "go",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".go",
				}},
			},
		}},
		{"StringGoCode", ki.Props{
			"label":       "Show Code",
			"desc":        "shows the Go-formatted initializer code, can be copy / pasted into program",
			"icon":        "go",
			"show-return": true,
		}},
		{"sep-diffs", ki.BlankProp{}},
		{"DiffsWithin", ki.Props{
			"desc":        "reports where the same param path is being set to different values within this set (both within the same Sheet and betwen sheets)",
			"icon":        "search",
			"show-return": true,
		}},
	},
}

var SetsProps = ki.Props{
	"ToolBar": ki.PropSlice{
		{"SaveJSON", ki.Props{
			"label": "Save As...",
			"desc":  "save to JSON formatted file",
			"icon":  "file-save",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".params",
				}},
			},
		}},
		{"OpenJSON", ki.Props{
			"label": "Open...",
			"desc":  "open from JSON formatted file",
			"icon":  "file-open",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".params",
				}},
			},
		}},
		{"sep-gocode", ki.BlankProp{}},
		{"SaveGoCode", ki.Props{
			"label": "Save Code As...",
			"desc":  "save to Go-formatted initializer code in file",
			"icon":  "go",
			"Args": ki.PropSlice{
				{"File Name", ki.Props{
					"ext": ".go",
				}},
			},
		}},
		{"StringGoCode", ki.Props{
			"label":       "Show Code",
			"desc":        "shows the Go-formatted initializer code, can be copy / pasted into program",
			"icon":        "go",
			"show-return": true,
		}},
		{"sep-diffs", ki.BlankProp{}},
		{"DiffsAll", ki.Props{
			"desc":        "between all sets, reports where the same param path is being set to different values",
			"icon":        "search",
			"show-return": true,
		}},
		{"DiffsFirst", ki.Props{
			"desc":        "between first set (e.g., the Base set) and rest of sets, reports where the same param path is being set to different values",
			"icon":        "search",
			"show-return": true,
		}},
		{"DiffsWithin", ki.Props{
			"desc":        "reports all the cases where the same param path is being set to different values within different sheets in given set",
			"icon":        "search",
			"show-return": true,
			"Args": ki.PropSlice{
				{"Set Name", ki.Props{}},
			},
		}},
	},
}
