// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/golang-framework/mvc/modules/crypto"
	"github.com/golang-framework/mvc/storage"
	s "src/storage"
)

type CryptoService struct {
}

func NewCryptoService() *CryptoService {
	return &CryptoService{}
}

func (srv *CryptoService) MD5(d string) *storage.Tpl {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	cry := crypto.New()
	cry.Mode = storage.Common //加密类型 common & hmac & rsa
	cry.D = []interface{}{
		storage.Md5, //加密方式
		d,           //加密字符串
	}

	o, errCryptoMD5 := cry.Engine()
	if errCryptoMD5 != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errCryptoMD5.Error()

		return res
	}

	res.Res = &storage.Y{
		"MD5": o,
	}

	return res
}

func (srv *CryptoService) SHA1(d string) *storage.Tpl {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	cry := crypto.New()
	cry.Mode = storage.Common //加密类型 common & hmac & rsa
	cry.D = []interface{}{
		storage.Sha1, //加密方式
		d,            //加密字符串
	}

	o, errCryptoSHA1 := cry.Engine()
	if errCryptoSHA1 != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errCryptoSHA1.Error()

		return res
	}

	res.Res = &storage.Y{
		"SHA1": o,
	}

	return res
}

func (srv *CryptoService) SHA256(d string) *storage.Tpl {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	cry := crypto.New()
	cry.Mode = storage.Common //加密类型 common & hmac & rsa
	cry.D = []interface{}{
		storage.Sha1, //加密方式
		d,            //加密字符串
	}

	o, errCryptoSHA256 := cry.Engine()
	if errCryptoSHA256 != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errCryptoSHA256.Error()

		return res
	}

	res.Res = &storage.Y{
		"SHA1": o,
	}

	return res
}

func (srv *CryptoService) HMAC(d string) *storage.Tpl {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	cry := crypto.New()
	cry.Mode = storage.Hmac //加密类型 common & hmac & rsa
	cry.D = []interface{}{
		storage.Md5,     //加密方式
		"salt_o123456o", //密钥
		d,               //加密字符串
	}

	o, errCryptoHMAC := cry.Engine()
	if errCryptoHMAC != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errCryptoHMAC.Error()

		return res
	}

	res.Res = &storage.Y{
		"HMAC": o,
	}

	return res
}

func (srv *CryptoService) AesEncode(d string) *storage.Tpl {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	cry := crypto.New()
	cry.Mode = storage.Aes
	cry.D = []interface{}{
		"o123456o", //  密钥
	}

	engine, errAesEngine := cry.Engine()
	if errAesEngine != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errAesEngine.Error()

		return res
	}

	aesEncode, errAesEncode := engine.(*crypto.Aes).Encrypt(d)
	if errAesEncode != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errAesEncode.Error()

		return res
	}

	res.Res = &storage.Y{
		"AES_Encode": aesEncode,
	}

	return res
}

func (srv *CryptoService) AesDecode(d string) *storage.Tpl {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	cry := crypto.New()
	cry.Mode = storage.Aes
	cry.D = []interface{}{
		"o123456o", //  密钥
	}

	engine, errAesEngine := cry.Engine()
	if errAesEngine != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errAesEngine.Error()

		return res
	}

	aesEncode, errAesEncode := engine.(*crypto.Aes).Decrypt(d)
	if errAesEncode != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errAesEncode.Error()

		return res
	}

	res.Res = &storage.Y{
		"AES_Decode": aesEncode,
	}

	return res
}
