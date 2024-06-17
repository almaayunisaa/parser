package main

import "fmt"

func S(x string) bool {
	if len(x) >= 1 && x[0] == 'k' {
		if len(x) >= 2 && x[1] == 'i' {
			if len(x) >= 3 && x[2] == 't' {
				if len(x) >= 4 && x[3] == 'a' {
					return true
				}
			}
		} else if len(x) >= 2 && x[1] == 'a' {
			if len(x) >= 3 && x[2] == 'u' {
				return true
			}
		}
	} else if len(x) >= 1 && x[0] == 'd' {
		if len(x) >= 2 && x[1] == 'i' {
			if len(x) >= 3 && x[2] == 'a' {
				return true
			}
		}
	} else if len(x) >= 1 && x[0] == 's' {
		if len(x) >= 2 && x[1] == 'a' {
			if len(x) >= 3 && x[2] == 'y' {
				if len(x) >= 4 && x[3] == 'a' {
					return true
				}
			}
		}
	} else if len(x) >= 1 && x[0] == 'a' {
		if len(x) >= 2 && x[1] == 'k' {
			if len(x) >= 3 && x[2] == 'u' {
				return true
			}
		}
	}
	return false
}

func P(x string) bool {
	if len(x) >= 3 && x[0] == 'm' && x[1] == 'e' && x[2] == 'm' {
		if len(x) >= 4 && x[3] == 'a' {
			if len(x) >= 5 && x[4] == 's' {
				if len(x) >= 6 && x[5] == 'a' {
					if len(x) >= 7 && x[6] == 'k' {
						return true
					}
				}
			} else if len(x) >= 5 && x[4] == 'k' {
				if len(x) >= 6 && x[5] == 'a' {
					if len(x) >= 7 && x[6] == 'n' {
						return true
					}
				}
			}
		} else if len(x) >= 4 && x[3] == 'b' {
			if len(x) >= 5 && x[4] == 'e' {
				if len(x) >= 6 && (x[5] == 'l' || x[5] == 'r') {
					if len(x) >= 7 && x[6] == 'i' {
						return true
					}
				}
			} else if len(x) >= 5 && x[4] == 'u' {
				if len(x) >= 6 && x[5] == 'a' {
					if len(x) >= 7 && x[6] == 't' {
						return true
					}
				}
			}
		}
	}
	return false
}

func O(x string) bool {
	if len(x) >= 1 && x[0] == 's' {
		if len(x) >= 2 && x[1] == 'o' {
			if len(x) >= 3 && x[2] == 'p' {
				return true
			}
		}
	} else if len(x) >= 1 && x[0] == 'r' {
		if len(x) >= 2 && x[1] == 'o' {
			if len(x) >= 3 && x[2] == 't' {
				if len(x) >= 4 && x[3] == 'i' {
					return true
				}
			}
		}
	} else if len(x) >= 1 && x[0] == 'k' {
		if len(x) >= 2 && x[1] == 'u' {
			if len(x) >= 3 && x[2] == 'e' {
				return true
			}
		}
	} else if len(x) >= 1 && x[0] == 'm' {
		if len(x) >= 2 && x[1] == 'i' {
			if len(x) >= 3 && x[2] == 'e' {
				return true
			}
		}
	} else if len(x) >= 1 && x[0] == 'b' {
		if len(x) >= 2 && x[1] == 'a' {
			if len(x) >= 3 && x[2] == 'k' {
				if len(x) >= 4 && x[3] == 's' {
					if len(x) >= 5 && x[4] == 'o' {
						return true
					}
				}
			}
		}
	}
	return false
}

func K(x string) bool {
	if len(x) >= 2 && x[0] == 'd' && x[1] == 'i' {
		if len(x) >= 3 && x[2] == 't' {
			if len(x) >= 4 && x[3] == 'o' {
				if len(x) >= 5 && x[4] == 'k' {
					if len(x) >= 6 && x[5] == 'o' {
						return true
					}
				}
			}
		} else if len(x) >= 3 && x[2] == 'k' {
			if len(x) >= 4 && x[3] == 'a' {
				if len(x) >= 5 && x[4] == 'f' {
					if len(x) >= 6 && x[5] == 'e' {
						return true
					}
				}
			} else if len(x) >= 4 && x[3] == 'o' {
				if len(x) >= 5 && x[4] == 's' {
					return true
				}
			}
		} else if len(x) >= 3 && x[2] == 'v' {
			if len(x) >= 4 && x[3] == 'i' {
				if len(x) >= 5 && x[4] == 'l' {
					if len(x) >= 6 && x[5] == 'a' {
						return true
					}
				}
			}
		} else if len(x) >= 3 && x[2] == 'm' {
			if len(x) >= 4 && x[3] == 'a' {
				if len(x) >= 5 && x[4] == 'l' {
					return true
				}
			}
		}
	}
	return false
}

func token(x string) byte {
	if S(x) {
		return 'S'
	} else if P(x) {
		return 'P'
	} else if O(x) {
		return 'O'
	} else if K(x) {
		return 'K'
	}
	return 'X'
}

type stack [100]string

func push(s *stack, x string, n *int) {
	s[*n] = x
	*n = *n + 1
}

func pop(s *stack, n *int) string {
	var hasil string
	hasil = s[*n-1]
	*n = *n - 1
	return hasil
}

func top(s stack, n int) string {
	return s[n-1]
}

func parserSPOK(x [100]byte, xN int) bool {
	var s stack
	var n int
	var i int = 0
	push(&s, "#", &n)
	push(&s, "A", &n)
	for top(s, n) != "#" && i < xN {
		if top(s, n) == "A" {
			pop(&s, &n)
			push(&s, "K", &n)
			push(&s, "O", &n)
			push(&s, "P", &n)
			push(&s, "S", &n)
		} else if top(s, n) == "S" {
			if x[i] != 'S' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "P" {
			if x[i] != 'P' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "O" {
			if x[i] != 'O' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "K" {
			if x[i] != 'K' {
				return false
			}
			i++
			pop(&s, &n)
		}
	}
	if pop(&s, &n) == "#" {
		return true
	}
	return false
}

func parserSPO(x [100]byte, xN int) bool {
	var s stack
	var n int
	var i int = 0
	push(&s, "#", &n)
	push(&s, "A", &n)
	for top(s, n) != "#" && i < xN {
		if top(s, n) == "A" {
			pop(&s, &n)
			push(&s, "O", &n)
			push(&s, "P", &n)
			push(&s, "S", &n)
		} else if top(s, n) == "S" {
			if x[i] != 'S' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "P" {
			if x[i] != 'P' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "O" {
			if x[i] != 'O' {
				return false
			}
			i++
			pop(&s, &n)
		}
	}
	if pop(&s, &n) == "#" {
		return true
	}
	return false
}

func parserSPK(x [100]byte, xN int) bool {
	var s stack
	var n int
	var i int = 0
	push(&s, "#", &n)
	push(&s, "A", &n)
	for top(s, n) != "#" && i < xN {
		if top(s, n) == "A" {
			pop(&s, &n)
			push(&s, "K", &n)
			push(&s, "P", &n)
			push(&s, "S", &n)
		} else if top(s, n) == "S" {
			if x[i] != 'S' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "P" {
			if x[i] != 'P' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "K" {
			if x[i] != 'K' {
				return false
			}
			i++
			pop(&s, &n)
		}
	}
	if pop(&s, &n) == "#" {
		return true
	}
	return false
}

func parserSP(x [100]byte, xN int) bool {
	var s stack
	var n int
	var i int = 0
	push(&s, "#", &n)
	push(&s, "A", &n)
	for top(s, n) != "#" && i < xN {
		if top(s, n) == "A" {
			pop(&s, &n)
			push(&s, "P", &n)
			push(&s, "S", &n)
		} else if top(s, n) == "S" {
			if x[i] != 'S' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "P" {
			if x[i] != 'P' {
				return false
			}
			i++
			pop(&s, &n)
		}
	}
	if pop(&s, &n) == "#" {
		return true
	}
	return false
}

func main() {
	// asumsi : ketik X untuk menyelesaikan input
	var kata string
	var himpKata [100]string
	var parse [100]byte
	var nParse int = 0
	var n int = 0
	var i int
	fmt.Scan(&kata)
	for kata != "X" {
		himpKata[n] = kata
		n++
		fmt.Scan(&kata)
	}
	for i = 0; i < n; i++ {
		if S(himpKata[i]) {
			parse[nParse] = 'S'
			nParse++
		} else if P(himpKata[i]) {
			parse[nParse] = 'P'
			nParse++
		} else if O(himpKata[i]) {
			parse[nParse] = 'O'
			nParse++
		} else if K(himpKata[i]) {
			parse[nParse] = 'K'
			nParse++
		}
	}

	fmt.Println("Apakah kalimat tersebut memenuhi tata bahasa Indonesia?")
	if nParse != 0 && (parserSPOK(parse, nParse) || parserSPO(parse, nParse) || parserSPK(parse, nParse) || parserSP(parse, nParse)) {
		fmt.Println("Iya")
	} else {
		fmt.Println("Tidak")
	}
}
