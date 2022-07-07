package utility

func CompareVersion(version1 string, version2 string) int {
	len1, len2, i, j := len(version1), len(version2), 0, 0
	for i < len1 || j < len2 {
		n1 := 0
		for i < len1 && '0' <= version1[i] && version1[i] <= '9' {
			n1 = n1*10 + int(version1[i]-'0')
			i++
		}
		n2 := 0
		for j < len2 && '0' <= version2[j] && version2[j] <= '9' {
			n2 = n2*10 + int(version2[j]-'0')
			j++
		}
		if n1 > n2 {
			return 1
		}
		if n1 < n2 {
			return -1
		}
		i, j = i+1, j+1
	}
	return 0
}
