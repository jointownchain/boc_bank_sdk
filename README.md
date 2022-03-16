# bocBankSDK 接口 使用说明

>  [XML 在线格式化](https://www.freeformatter.com/xml-formatter.html#ad-output)
> 
>  [XML 转 struct 在线工具](https://www.onlinetool.io/xmltogo/)

## 关键文件 api.go
包含所有接口

## 调用步骤
1. 构造 xml 请求头
`NewXMLRequestHead`
2. 构造 xml 请求体
`new(TrnB2e0009Rq)/或者别的`
3. 调用 HospitalPayment/或者别的 方法

## 接口添加说明
1. **添加 request 和 response 结构体**
    - 位于 Modules/request.go
    - 位于 Modules/response.go
      - **额外添加 CompanyFactoringRspTrans 类似的结构体供传递给 interface**
    - 拆分 struct
2. ~~添加 NewStruct 函数~~
    - ~~位于 ./request.go 中~~
    - ~~添加构造 request 的函数, 返回 类似 *Modules.TrnB2e0477Rq 的指针~~
3. **添加 response case**
    - 位于 ./response.go 中
    - 在 ParseResponseXML **方法** 中添加新的 case
        - opts 的 const 定义在 Modules/constant.go
``` go
	case Modules.BOCB2E_OPTS_RAISE_REQUEST:
		trans = &Modules.RaiseRequestRspTrans{}
```
4. 添加 api.go 中暴露的方法
``` go
// RaiseRequest 提额申请 b2e0057
func RaiseRequest(bh *Modules.Head, trnRq *Modules.TrnB2e0057Rq) (*Modules.RaiseRequestRspTrans, error) {
	xmlReqBytes, err := NewXMLRequest(bh, trnRq)
	if err != nil {
		return nil, err
	}

	in, err := doRequest(xmlReqBytes, Modules.BOCB2E_OPTS_RAISE_REQUEST)
	if err != nil {
		return nil, err
	}
	rsp, ok := in.(*Modules.RaiseRequestRspTrans)
	if !ok {
		return nil, errors.New("interface revert error")
	}
	// 错误响应码返回
	if rsp.TrnB2e0057Rs.Status.Rspcod != Modules.BOCB2E_RESPONSE_CODE_OK {
		return nil, errors.Errorf("RaiseRequest TRN error: code : %s, msg: %s",
			rsp.TrnB2e0057Rs.Status.Rspcod, rsp.TrnB2e0057Rs.Status.Rspmsg)
	}
	if rsp.TrnB2e0057Rs.B2e0057Rs.Status.Rspcod != Modules.BOCB2E_RESPONSE_CODE_OK {
		return nil, errors.Errorf("RaiseRequest B2E error: code : %s, msg: %s",
			rsp.TrnB2e0057Rs.B2e0057Rs.Status.Rspcod, rsp.TrnB2e0057Rs.B2e0057Rs.Status.Rspmsg)
	}
	return rsp, nil
}
```

## 目录结构说明
```shell
.
├── api.go # 主接口文件
├── Modules # struct 定义
│   ├── bocB2eRoot.go # 上层结构体定义
│   ├── constant.go 
│   ├── error.go # 错误返回 xml 定义
│   ├── head.go # xml head 定义(通用)
│   ├── request.go # 请求类型 xml 定义
│   └── response.go # 返回类型 xml 定义
├── go.mod
├── go.sum
├── javaDemo # java 示例文件夹
├── main_test.go
├── modules_test.go
├── README.md
├── request.go # 请求发起逻辑
├── response.go # 接收逻辑
├── response_test.go
├── rspError # error 接收中间件
│   └── error.go
├── util.go # 工具包
└── xmls # xml 定义源文件
    ├── b2e0009-req.xml
    ├── b2e0009-rsp.xml
    ├── b2e0477-req.xml
    ├── b2e0477-rsp.xml
    ├── b2e0481-req.xml
    └── b2e0481-rsp.xml

```

[]: https://www.freeformatter.com/xml-formatter.html#ad-output