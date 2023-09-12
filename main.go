package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type File struct {
	Name       string `json:"name"`
	Expression Expression
	Location   struct {
		Start    int    `json:"start"`
		End      int    `json:"end"`
		Filename string `json:"filename"`
	} `json:"location"`
}

type Expression struct {
	Kind string `json:"kind"`
	Name struct {
		Text     string `json:"text"`
		Location struct {
			Start    int    `json:"start"`
			End      int    `json:"end"`
			Filename string `json:"filename"`
		} `json:"location"`
	} `json:"name"`
	Value Value
	Next  struct {
		Kind  string `json:"kind"`
		Value struct {
			Kind   string `json:"kind"`
			Callee struct {
				Kind     string `json:"kind"`
				Text     string `json:"text"`
				Location struct {
					Start    int    `json:"start"`
					End      int    `json:"end"`
					Filename string `json:"filename"`
				} `json:"location"`
			} `json:"callee"`
			Arguments []struct {
				Kind     string `json:"kind"`
				Value    int    `json:"value"`
				Location struct {
					Start    int    `json:"start"`
					End      int    `json:"end"`
					Filename string `json:"filename"`
				} `json:"location"`
			} `json:"arguments"`
			Location struct {
				Start    int    `json:"start"`
				End      int    `json:"end"`
				Filename string `json:"filename"`
			} `json:"location"`
		} `json:"value"`
		Location struct {
			Start    int    `json:"start"`
			End      int    `json:"end"`
			Filename string `json:"filename"`
		} `json:"location"`
	} `json:"next"`
	Location struct {
		Start    int    `json:"start"`
		End      int    `json:"end"`
		Filename string `json:"filename"`
	} `json:"location"`
}

type Value struct {
	Kind       string `json:"kind"`
	Parameters []struct {
		Text     string `json:"text"`
		Location struct {
			Start    int    `json:"start"`
			End      int    `json:"end"`
			Filename string `json:"filename"`
		} `json:"location"`
	} `json:"parameters"`
	Value    ValueValue
	Location struct {
		Start    int    `json:"start"`
		End      int    `json:"end"`
		Filename string `json:"filename"`
	} `json:"location"`
}

type ValueValue struct {
	Kind      string `json:"kind"`
	Condition struct {
		Kind string `json:"kind"`
		LHS  struct {
			Kind     string `json:"kind"`
			Text     string `json:"text"`
			Location struct {
				Start    int    `json:"start"`
				End      int    `json:"end"`
				Filename string `json:"filename"`
			} `json:"location"`
		} `json:"lhs"`
		Op  string `json:"op"`
		RHS struct {
			Kind     string `json:"kind"`
			Value    int    `json:"value"`
			Location struct {
				Start    int    `json:"start"`
				End      int    `json:"end"`
				Filename string `json:"filename"`
			} `json:"location"`
		} `json:"rhs"`
		Location struct {
			Start    int    `json:"start"`
			End      int    `json:"end"`
			Filename string `json:"filename"`
		} `json:"location"`
	} `json:"condition"`
	Then struct {
		Kind     string `json:"kind"`
		Text     string `json:"text"`
		Location struct {
			Start    int    `json:"start"`
			End      int    `json:"end"`
			Filename string `json:"filename"`
		} `json:"location"`
	} `json:"then"`
	Otherwise struct {
		Kind string `json:"kind"`
		LHS  struct {
			Kind   string `json:"kind"`
			Callee struct {
				Kind     string `json:"kind"`
				Text     string `json:"text"`
				Location struct {
					Start    int    `json:"start"`
					End      int    `json:"end"`
					Filename string `json:"filename"`
				} `json:"location"`
			} `json:"callee"`
			Arguments []struct {
				Kind string `json:"kind"`
				LHS  struct {
					Kind     string `json:"kind"`
					Text     string `json:"text"`
					Location struct {
						Start    int    `json:"start"`
						End      int    `json:"end"`
						Filename string `json:"filename"`
					} `json:"location"`
				} `json:"lhs"`
				Op  string `json:"op"`
				RHS struct {
					Kind     string `json:"kind"`
					Value    int    `json:"value"`
					Location struct {
						Start    int    `json:"start"`
						End      int    `json:"end"`
						Filename string `json:"filename"`
					} `json:"location"`
				} `json:"rhs"`
				Location struct {
					Start    int    `json:"start"`
					End      int    `json:"end"`
					Filename string `json:"filename"`
				} `json:"location"`
			} `json:"arguments"`
			Location struct {
				Start    int    `json:"start"`
				End      int    `json:"end"`
				Filename string `json:"filename"`
			} `json:"location"`
		} `json:"lhs"`
		Op  string `json:"op"`
		RHS struct {
			Kind   string `json:"kind"`
			Callee struct {
				Kind     string `json:"kind"`
				Text     string `json:"text"`
				Location struct {
					Start    int    `json:"start"`
					End      int    `json:"end"`
					Filename string `json:"filename"`
				} `json:"location"`
			} `json:"callee"`
			Arguments []struct {
				Kind string `json:"kind"`
				LHS  struct {
					Kind     string `json:"kind"`
					Text     string `json:"text"`
					Location struct {
						Start    int    `json:"start"`
						End      int    `json:"end"`
						Filename string `json:"filename"`
					} `json:"location"`
				} `json:"lhs"`
				Op  string `json:"op"`
				RHS struct {
					Kind     string `json:"kind"`
					Value    int    `json:"value"`
					Location struct {
						Start    int    `json:"start"`
						End      int    `json:"end"`
						Filename string `json:"filename"`
					} `json:"location"`
				} `json:"rhs"`
				Location struct {
					Start    int    `json:"start"`
					End      int    `json:"end"`
					Filename string `json:"filename"`
				} `json:"location"`
			} `json:"arguments"`
			Location struct {
				Start    int    `json:"start"`
				End      int    `json:"end"`
				Filename string `json:"filename"`
			} `json:"location"`
		} `json:"rhs"`
		Location struct {
			Start    int    `json:"start"`
			End      int    `json:"end"`
			Filename string `json:"filename"`
		} `json:"location"`
	} `json:"otherwise"`
	Location struct {
		Start    int    `json:"start"`
		End      int    `json:"end"`
		Filename string `json:"filename"`
	} `json:"location"`
}

func main() {
	file, err := os.ReadFile("./files/fb.json")
	if err != nil {
		panic(err)
	}

	// Read json file
	var f File
	err = json.Unmarshal(file, &f)
	if err != nil {
		panic(err)
	}

	// parse the ast file and print the result
	var code []string
	code = interpreter(f.Expression, code)

	for _, v := range code {
		fmt.Printf("%s\n", v)
	}

}

func interpreter(node Expression, code []string) []string {
	switch node.Kind {
	case "Let":
		code = append(code, "var")
		code = append(code, node.Name.Text)
		return getValueKind(node.Value, code)
	}
	return code
}

func getValueKind(node Value, code []string) []string {
	switch node.Kind {
	case "Function":
		if node.Parameters != nil {
			params := ""
			for i, v := range node.Parameters {
				if len(node.Parameters) > 1 && i < len(node.Parameters)-1 {
					params += v.Text + ","
				} else {
					params += v.Text
				}
			}
			code = append(code, fmt.Sprintf("func(%s)", params))
			return getValueValueKind(node.Value, code)
		}
	}
	return code
}

func getValueValueKind(node ValueValue, code []string) []string {
	switch node.Kind {
	case "If":
		if node.Condition.Kind == "Binary" {
			if node.Condition.Op == "Lt" {
				code = append(code, fmt.Sprintf("if %s < %d {", node.Condition.LHS.Text, node.Condition.RHS.Value))
				return code
			}
		}
	}
	return code
}
