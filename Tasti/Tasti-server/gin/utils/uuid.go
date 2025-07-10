package utils

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// GenerateV4 生成版本4的UUID（随机生成）
func GenerateV4() string {
	return uuid.New().String()
}

// GenerateV4WithoutHyphen 生成无连字符的版本4 UUID
func GenerateV4WithoutHyphen() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// GenerateV4WithPrefix 生成带前缀的版本4 UUID
func GenerateV4WithPrefix(prefix string) string {
	return fmt.Sprintf("%s%s", prefix, uuid.New().String())
}

// GenerateV1 生成版本1的UUID（基于时间和MAC地址）
func GenerateV1() (string, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// Parse 解析UUID字符串
func Parse(uuidStr string) (uuid.UUID, error) {
	return uuid.Parse(uuidStr)
}

// IsValid 检查字符串是否为有效的UUID格式
func IsValid(uuidStr string) bool {
	_, err := uuid.Parse(uuidStr)
	return err == nil
}
