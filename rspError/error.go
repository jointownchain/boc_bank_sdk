package rspError

import (
	"encoding/xml"

	"github.com/beevik/etree"
	"github.com/jointownchain/boc_bank_sdk/Modules"
	"github.com/pkg/errors"
)

// 前置判断是否为错误xml返回
func ParseError(xmlBytes []byte) (err error) {
	doc := etree.NewDocument()
	err = doc.ReadFromBytes(xmlBytes)
	if err != nil {
		return errors.Wrap(err, "ParseError")
	}
	e := doc.FindElement(Modules.BOCB2E_ERROR_XML_PATH)
	if e == nil {
		return nil
	}
	errFlag := false
	for _, child := range e.ChildElements() {
		if child.Tag == Modules.BOCB2E_ERROR_XML_KEY {
			errFlag = true
		}
	}
	if errFlag {
		be, err := unmarshalError(xmlBytes)
		if err != nil {
			return errors.Wrap(err, "ParseError")
		}
		return errors.New("code:" + be.Trans.TrnB2eerrorRs.Status.Rspcod + ", msg:" + be.Trans.TrnB2eerrorRs.Status.Rspmsg)
	}
	return nil
}

// 解析 error 类型的 xml
func unmarshalError(xmlBytes []byte) (be *Modules.Error, err error) {
	be = &Modules.Error{}
	err = xml.Unmarshal(xmlBytes, be)
	if err != nil {
		return nil, err
	}
	return be, err
}
