package model

import (
	errmsg "GGblog/internal/errormsg"
	"GGblog/internal/setting"
	"context"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: setting.QiniuConf.Bucket,
	}
	mac := qbox.NewMac(setting.QiniuConf.AccessKey, setting.QiniuConf.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Region:        &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := setting.QiniuConf.Server + ret.Key
	return url, errmsg.SUCCESS
}
