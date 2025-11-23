package piscine

func WeAreUnique(str1 , str2 string) int {
    setA := uniqueRuneSet(str1)
    setB := uniqueRuneSet(str2)

    count := 0

    for r := range setA {
        if _, exists := setB[r]; !exists {
            count++
        }
    }

    for r := range setB {
        if _, exists := setA[r]; !exists {
            count++
        }
    }

    return count
}


func uniqueRuneSet(s string) map[rune]struct{} {
    m := make(map[rune]struct{})
    for _, r := range s {
        m[r] = struct{}{}
    }
    return m
}


 