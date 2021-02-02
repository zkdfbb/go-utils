package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/intel-go/cpuid"
	"reflect"
	"runtime"
)

// GetFuncName return func name
func GetFuncName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// GetLocalIP return local ip
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("Can not find the client ip address")
}

// RandomStr random str
func RandomStr(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 97
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// RandomIP ip address
func RandomIP() string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

// Md5 for string
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// FileMd5 calc file md5
func FileMd5(file string) string {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return ""
	}
	h := md5.New()
	io.Copy(h, f)
	return hex.EncodeToString(h.Sum(nil))
}

// UUID unique id
func UUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

// CPUID unique cpu id
func CPUID() string {
	features := []string{}
	for i := uint64(0); i < 64; i++ {
		if cpuid.HasExtendedFeature(1 << i) {
			features = append(features, cpuid.ExtendedFeatureNames[1<<i])
		}
	}
	return Md5(strings.Join(features, " "))
}

// URLJoin join urls
func URLJoin(base, href string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return href
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		return href
	}
	return baseURL.ResolveReference(uri).String()
}

// Exists files
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsDir check
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile check
func IsFile(path string) bool {
	return Exists(path) && !IsDir(path)
}
