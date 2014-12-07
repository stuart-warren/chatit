package main

import (
    "crypto/md5"
    "encoding/hex"
    "code.google.com/p/go-uuid/uuid"
)

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func GetUUID() string {
    return uuid.New()
}
