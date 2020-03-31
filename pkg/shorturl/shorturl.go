package shorturl


import (
    "strings"
)

// 定义基本的配置
const (
    default_alphabet = "asdfghjklURLEncoderConfigqwertyui"
    default_block_size = uint(24)
    default_min_length = 5
)

// url encoder结构体
type UrlEncoder struct {
    alphabet string
    block_size uint
}

// 定义配置结构体
type EncoderConfig struct {
    alphabet string
    block_size uint
}

// 定义new url encoder
func NewUrlEncoder(config *EncoderConfig) *UrlEncoder {
    // 默认定义的配置
    alphabet := default_alphabet
    block_size := default_block_size
    // 如果设置alphabet和block_size则用传入的填充
    if config.alphabet != "" {
        alphabet = config.alphabet
    }

    if config.block_size != 0 {
                block_size= config.block_size
        }

        urlencoder := &UrlEncoder{
                alphabet: alphabet,
                block_size: block_size,
        }

        return urlencoder
}

// 0,1 check
func getBit(n uint64, pos uint) int {
        if (n & (1 << pos)) != 0 {
                return 1
        }
        return 0
}

func (encoder *UrlEncoder) encode(n uint64) uint64 {
        for i, j := uint(0), uint(encoder.block_size-1); i < j;i, j = i+1, j-1 {
                if getBit(n, i) != getBit(n, j) {
                        n ^= ((1 << i) | (1 << j))
                }
        }
        return n
}

func (encoder *UrlEncoder) enbase(x uint64) string {
        n := uint64(len(encoder.alphabet))
        result := []byte{}

        for {
                ch := encoder.alphabet[x%n]
                result = append(result, ch)
                x = x / n
                if x == 0 && len(result) >= default_min_length {
                        break
                }
        }
        // 反转结果集
        revResult := []byte{}
        for i := len(result)-1; i >= 0; i-- {
                revResult = append(revResult, result[i])
        }
        return string(revResult)
}

func (encoder *UrlEncoder) debase(x string) uint64 {
        n := uint64(len(encoder.alphabet))
        result := uint64(0)
        bits := []byte(x)

        for _, bitValue := range bits {
                result = result*n + uint64(strings.IndexByte(encoder.alphabet, bitValue))
        }
        return result
}

func (encoder *UrlEncoder) EncodeURL(n uint64) string {
        return encoder.enbase(encoder.encode(n))
}


func (encoder *UrlEncoder) DecodeURL(n string) uint64{
        return encoder.encode(encoder.debase(n))
}
