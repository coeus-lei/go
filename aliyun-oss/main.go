package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	Endpoint        = "xxxxxxx"
	AccessKeyId     = "xxxxxxx"
	AccessKeySecret = "xxxxxxx"
)

func main() {
	// fmt.Println("OSS Go SDK Version: ", oss.Version)
	str, _ := os.Getwd()
	_path := filepath.Join(str, `pyq_h`, `hotupdate`, `site`)
	fmt.Println(_path)
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}
	// lsRes, err := client.ListBuckets()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, bucket := range lsRes.Buckets {
	// 	fmt.Println("Buckets:", bucket.Name)
	// }
	bucket, err := client.Bucket("cocos-publishv1")
	if err != nil {
		fmt.Println(err)
	}

	// err = bucket.PutObjectFromFile("pyq_h", _path)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	lsRes, err := bucket.ListObjects()
	if err != nil {
		fmt.Println(err)
	}
	for _, object := range lsRes.Objects {
		fmt.Println("Objects:", object.Key)
	}

	// err = bucket.DeleteObject("pyq_h")
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
