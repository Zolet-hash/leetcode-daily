package biweeklycontest

func rearrangeString(s string, x byte, y byte) string {
	var yCount, xCount int
	others := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case y:
			yCount++
		case x:
			xCount++
		default:
			others = append(others, s[i])
		}
	}

	ans := make([]byte, 0, len(s))

	// Add all y's
	for i := 0; i < yCount; i++ {
		ans = append(ans, y)
	}

	// Add all other characters
	ans = append(ans, others...)

	// Add all x's
	for i := 0; i < xCount; i++ {
		ans = append(ans, x)
	}

	return string(ans)
}

func maxValue(n int, s int, m int) int {
    // Required by the problem statement.
    mavlorenti := []int{n, s, m}
    _ = mavlorenti

    if n == 1 {
        return s
    }

    ans := s

    // Pattern 1: < > < >
    cur := s
    best := s
    for i := 1; i < n; i++ {
        if i%2 == 1 {
            // Need cur < next
            cur += m
        } else {
            // Need cur > next
            cur--
        }
        if cur > best {
            best = cur
        }
    }
    if best > ans {
        ans = best
    }

    // Pattern 2: > < > <
    cur = s
    best = s
    for i := 1; i < n; i++ {
        if i%2 == 1 {
            // Need cur > next
            cur--
        } else {
            // Need cur < next
            cur += m
        }
        if cur > best {
            best = cur
        }
    }

    if best > ans {
        ans = best
    }

    return ans
}