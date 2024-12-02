package xyz_hash

func CsharpStringHashV1(str string) int64 {
	var num1, num2 int32 = 5381, 5381

	length := len(str)
	for i, nexti := 0, 0; i < length; i += 2 {
		num1 = (num1 << 5) + num1 ^ int32(str[i])
		nexti = i + 1
		if nexti != length {
			num2 = (num2 << 5) + num2 ^ int32(str[nexti])
		}
	}

	return int64(num1+num2*1566083941) & 0xFFFFFFFF
}

func CsharpStringHashV2(str string) int64 {
	var num1 int32 = 352654597
	var num2 int32 = num1

	var length int
	for length = len(str); length > 2; length -= 4 {
		num1 = (num1 << 5) + num1 + (num1 >> 27) ^ int32(str[0])
		num2 = (num2 << 5) + num2 + (num2 >> 27) ^ int32(str[1])
		str = str[2:]
	}
	if length > 0 {
		num1 = (num1 << 5) + num1 + (num1 >> 27) ^ int32(str[0])
	}
	return int64(num1+num2*1566083941) & 0xFFFFFFFF
}
