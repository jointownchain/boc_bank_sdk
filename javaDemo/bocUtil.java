package com.yizai.banks.boc;

import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

import org.jdom.Document;

import com.yizai.HandleBankUtility;
import com.yizai.HandleException;
import com.yizai.UtilConfigReader;
import com.yizai.bcp.boc.boc001.BOCImpl;
import com.yizai.structure.BaseEnumBankCode;
import com.yizai.structure.BaseEnumIsoCurrency;
import com.yizai.utils.UtilXML;

public class bocUtil
{
	private static Map<BaseEnumIsoCurrency, String> currencyCodeMap = new HashMap<BaseEnumIsoCurrency, String>();
	private static Map<String, String> docTypeMap = new HashMap<String, String>();

	public static Map<BaseEnumIsoCurrency, String> getCurrencyCodeMap()
	{
		if (currencyCodeMap.size() == 0)
		{
//			001	CNY
//			012	GBP
//			013	HKD
//			014	USD
//			015	CHF
//			018	SGD
//			021	SEK
//			022	DKK
//			023	NOK
//			027	JPY
//			028	CAD
//			029	AUD
//			038	EUR
//			072	RUR
//			081	MOP
//			082	PHP
//			084	THB
//			087	NZD
//			101	KZT
			currencyCodeMap.put(BaseEnumIsoCurrency.RMB, "001");
			currencyCodeMap.put(BaseEnumIsoCurrency.CNY, "001");
			currencyCodeMap.put(BaseEnumIsoCurrency.GBP, "012");
			currencyCodeMap.put(BaseEnumIsoCurrency.HKD, "013");
			currencyCodeMap.put(BaseEnumIsoCurrency.USD, "014");
			currencyCodeMap.put(BaseEnumIsoCurrency.CHF, "015");
			currencyCodeMap.put(BaseEnumIsoCurrency.SGD, "018");
			currencyCodeMap.put(BaseEnumIsoCurrency.SEK, "021");
			currencyCodeMap.put(BaseEnumIsoCurrency.DKK, "022");
			currencyCodeMap.put(BaseEnumIsoCurrency.NOK, "023");
			currencyCodeMap.put(BaseEnumIsoCurrency.JPY, "027");
			currencyCodeMap.put(BaseEnumIsoCurrency.CAD, "028");
			currencyCodeMap.put(BaseEnumIsoCurrency.AUD, "029");
			currencyCodeMap.put(BaseEnumIsoCurrency.EUR, "038");
			currencyCodeMap.put(BaseEnumIsoCurrency.RUR, "072");
			currencyCodeMap.put(BaseEnumIsoCurrency.MOP, "081");
			currencyCodeMap.put(BaseEnumIsoCurrency.PHP, "082");
			currencyCodeMap.put(BaseEnumIsoCurrency.THB, "084");
			currencyCodeMap.put(BaseEnumIsoCurrency.NZD, "087");
			currencyCodeMap.put(BaseEnumIsoCurrency.KZT, "101");
		}
		return currencyCodeMap;
	}
	
	public static void setHeadForRequest(Document document, BaseEnumBankCode bankCode) throws Exception
	{
		String number17 = new SimpleDateFormat("yyyyMMddHHmmssSSS").format(new Date());
		String packetId = number17.substring(5);
		String termid = UtilConfigReader.getBcpConfig(bankCode.getCode(), "terminalId");
		String custid = UtilConfigReader.getBcpConfig(bankCode.getCode(), "customerId");
		String operator = UtilConfigReader.getBcpConfig(bankCode.getCode(), "userId");
		String token = BOCImpl.loginToken;//bocLogin.loginToken;
		String headPath = "/" + document.getRootElement().getName() + "/head";
		UtilXML.setText(document, headPath + "/termid", termid);
		UtilXML.setText(document, headPath + "/trnid", packetId);
		UtilXML.setText(document, headPath + "/custid", custid);
		UtilXML.setText(document, headPath + "/cusopr", operator);
		UtilXML.setText(document, headPath + "/token", token);
	}

	public static String setCityCodeForBOC(String bank_key) throws HandleException
	{
		int shortRegionCode = HandleBankUtility.getBankRegionCode(bank_key, 5);
		int fullRegionCode = HandleBankUtility.getBankRegionCode(bank_key, 7);
		String bocInternalKey = "";
		switch (shortRegionCode)
		{
		// 北京
		case 10:
			bocInternalKey = "40142";
			break;
		// 天津
		case 11:
			bocInternalKey = "40202";
			break;
		// 河北
		case 12:
			bocInternalKey = "40740";
			break;
		// 河北
		case 13:
			bocInternalKey = "40740";
			break;
		// /河北
		case 14:
			bocInternalKey = "40740";
			break;
		// 山西
		case 16:
			bocInternalKey = "41041";
			break;
		// 山西
		case 17:
			bocInternalKey = "41041";
			break;
		// 山西
		case 18:
			bocInternalKey = "41041";
			break;
		// 内蒙古
		case 19:
			bocInternalKey = "41405";
			break;
		// 内蒙古
		case 20:
			bocInternalKey = "41405";
			break;
		// 辽宁
		case 22:
			bocInternalKey = "41785";
			break;
		// 辽宁
		case 23:
			bocInternalKey = "41785";
			break;
		// 吉林
		case 24:
			bocInternalKey = "42208";
			break;
		// 吉林
		case 25:
			bocInternalKey = "42208";
			break;
		// 黑龙江
		case 26:
			bocInternalKey = "42465";
			break;
		// 黑龙江
		case 27:
			bocInternalKey = "42465";
			break;
		// 上海
		case 29:
			bocInternalKey = "40303";
			break;
		// 江苏
		case 30:
			bocInternalKey = "44433";
			break;
		// 江苏
		case 31:
			bocInternalKey = "44433";
			break;
		// 浙江
		case 33:
			bocInternalKey = "45129";
			break;
		// 浙江
		case 34:
			bocInternalKey = "45129";
			break;
		// 安徽
		case 36:
			bocInternalKey = "44899";
			break;
		// 安徽
		case 37:
			bocInternalKey = "44899";
			break;
		// 福建
		case 39:
			bocInternalKey = "45481";
			break;
		// 福建
		case 40:
			bocInternalKey = "45481";
			break;
		// 江西
		case 42:
			bocInternalKey = "47370";
			break;
		// 江西
		case 43:
			bocInternalKey = "47370";
			break;
		// 山东
		case 45:
			bocInternalKey = "43810";
			break;
		// 山东
		case 46:
			bocInternalKey = "43810";
			break;
		// 山东
		case 47:
			bocInternalKey = "43810";
			break;
		// 河南
		case 49:
			bocInternalKey = "46243";
			break;
		// 河南
		case 50:
			bocInternalKey = "46243";
			break;
		// 河南
		case 51:
			bocInternalKey = "46243";
			break;
		// 湖北
		case 52:
			bocInternalKey = "46405";
			break;
		// 湖北
		case 53:
			bocInternalKey = "46405";
			break;
		// 湖北
		case 54:
			bocInternalKey = "46405";
			break;
		// 湖南
		case 55:
			bocInternalKey = "46955";
			break;
		// 湖南
		case 56:
			bocInternalKey = "46955";
			break;
		// 广东
		case 58:
			if (fullRegionCode == 5840)
			{
				// 深证
				bocInternalKey = "47669";
			} else
			{
				// 其余的广东省城市
				bocInternalKey = "47504";
			}
			break;
		// 广东
		case 59:
			bocInternalKey = "47504";
			break;
		// 广东
		case 60:
			bocInternalKey = "47504";
			break;
		// 广西
		case 61:
			bocInternalKey = "48051";
			break;
		// 广西
		case 62:
			bocInternalKey = "48051";
			break;
		// 广西
		case 63:
			bocInternalKey = "48051";
			break;
		// 海南
		case 64:
			bocInternalKey = "47806";
			break;
		// 四川
		case 65:
			bocInternalKey = "48631";
			break;
		// 四川
		case 66:
			bocInternalKey = "48631";
			break;
		// 四川
		case 67:
			bocInternalKey = "48631";
			break;
		// 四川
		case 68:
			if (fullRegionCode >= 6870 && fullRegionCode <= 6875)
			{
				// 重庆
				bocInternalKey = "48642";
			} else
			{
				// 其余的四川省城市
				bocInternalKey = "48631";
			}
			break;
		// 贵州
		case 70:
			bocInternalKey = "48882";
			break;
		// 贵州
		case 71:
			bocInternalKey = "48882";
			break;
		// 云南
		case 73:
			bocInternalKey = "49146";
			break;
		// 云南
		case 74:
			bocInternalKey = "49146";
			break;
		// 云南
		case 75:
			bocInternalKey = "49146";
			break;
		// 西藏
		case 77:
			bocInternalKey = "40600";
			break;
		// 西藏
		case 78:
			bocInternalKey = "40600";
			break;
		// 陕西
		case 79:
			bocInternalKey = "43016";
			break;
		// 陕西
		case 80:
			bocInternalKey = "43016";
			break;
		// 甘肃
		case 82:
			bocInternalKey = "43251";
			break;
		// 甘肃
		case 83:
			bocInternalKey = "43251";
			break;
		// 青海
		case 85:
			bocInternalKey = "43469";
			break;
		// 宁夏
		case 87:
			bocInternalKey = "43347";
			break;
		// 新疆
		case 88:
			bocInternalKey = "43600";
			break;
		// 新疆
		case 89:
			bocInternalKey = "43600";
			break;
		// 新疆
		case 90:
			bocInternalKey = "43600";
			break;
		}
		// 未分配出去的代码
		// 省市代码(40)
		// 71 台湾省
		// 81 营业部
		// 97 香港
		// 98 澳门
		// 99 总行

		if (bocInternalKey.equals(""))
		{
			throw new HandleException("[BCP系统消息]" + "无法根据联行号解析出中 行的内部联行号，BCP接受到的值为" + bank_key);
		} else
		{
			return bocInternalKey;
		}
	}
	
	private static void initDocTypeMap()
	{
		docTypeMap.put("1011", "运通旅行支票(50美元)");
		docTypeMap.put("1012", "运通旅行支票(100美元)");
		docTypeMap.put("1013", "运通旅行支票(500美元)");
		docTypeMap.put("1020", "运通旅行支票(50欧元)");
		docTypeMap.put("1021", "运通旅行支票(100欧元)");
		docTypeMap.put("1022", "运通旅行支票(200欧元)");
		docTypeMap.put("1023", "运通旅行支票(500欧元)");
		docTypeMap.put("1031", "运通旅行支票(50英镑)");
		docTypeMap.put("1032", "运通旅行支票(100英镑)");
		docTypeMap.put("1033", "运通旅行支票(200英镑)");
		docTypeMap.put("1040", "运通旅行支票(20澳元)");
		docTypeMap.put("1041", "运通旅行支票(50澳元)");
		docTypeMap.put("1042", "运通旅行支票(100澳元)");
		docTypeMap.put("1043", "运通旅行支票(200澳元)");
		docTypeMap.put("1050", "运通旅行支票(10000日元)");
		docTypeMap.put("1051", "运通旅行支票(20000日元)");
		docTypeMap.put("1052", "运通旅行支票(50000日元)");
		docTypeMap.put("1060", "运通旅行支票(20加元)");
		docTypeMap.put("1061", "运通旅行支票(50加元)");
		docTypeMap.put("1062", "运通旅行支票(100加元)");
		docTypeMap.put("1063", "运通旅行支票(500加元)");
		docTypeMap.put("1101", "电子旅行支票（美元）");
		docTypeMap.put("1102", "电子旅行支票（欧元）");
		docTypeMap.put("1103", "电子旅行支票（英镑）");
		docTypeMap.put("2001", "个人定期存单(礼仪存单)");
		docTypeMap.put("2002", "单位定期存单");
		docTypeMap.put("2201", "活期一本通存折");
		docTypeMap.put("2202", "定期一本通存折");
		docTypeMap.put("2203", "活期一本通存折(成长账户)");
		docTypeMap.put("2204", "定期一本通存折(成长账户)");
		docTypeMap.put("2205", "活期一本通存折(财富管理)");
		docTypeMap.put("2206", "定期一本通存折(财富管理)");
		docTypeMap.put("3001", "清分机支票(转账)");
		docTypeMap.put("3002", "现金支票");
		docTypeMap.put("3003", "转账支票");
		docTypeMap.put("3004", "清分机支票(普通)");
		docTypeMap.put("3005", "驻华机构人民币支票");
		docTypeMap.put("3006", "驻华机构外币支票");
		docTypeMap.put("3201", "全国银行汇票");
		docTypeMap.put("3202", "华东三省一市银行汇票");
		docTypeMap.put("3203", "银行承兑汇票");
		docTypeMap.put("3204", "商业承兑汇票");
		docTypeMap.put("3205", "纽约中行美元磁性汇票");
		docTypeMap.put("3206", "伦敦中行英镑磁性汇票");
		docTypeMap.put("3207", "自由格式汇票");
		docTypeMap.put("3401", "清分机本票");
		docTypeMap.put("3402", "非清分机本票");
		docTypeMap.put("4001", "借记卡");
		docTypeMap.put("7002", "凭证式国债存款凭证");
		docTypeMap.put("7003", "存款证明");
		docTypeMap.put("7004", "携带外汇出境许可证");
		docTypeMap.put("7005", "企业存款证明");
		docTypeMap.put("7006", "单位定期存款开户证实书");
		docTypeMap.put("7007", "长城借记卡对账簿(活期类交易)");
		docTypeMap.put("7008", "长城借记卡对账簿(定期类交易)");
		docTypeMap.put("3208", "巴黎中行欧元磁性汇票");
		docTypeMap.put("9001", "支取凭证");
		docTypeMap.put("X001", "电子商业汇票");
		docTypeMap.put("A001", "商业承兑汇票(非重空)");
		docTypeMap.put("A002", "存单/存折");
		docTypeMap.put("A003", "光票托收");
		docTypeMap.put("A004", "验单付款");
		docTypeMap.put("A005", "验货付款");
		docTypeMap.put("A006", "国内信用凭证");
		docTypeMap.put("A007", "非我行转账支票");
		docTypeMap.put("A008", "非我行全国银行汇票");
		docTypeMap.put("A009", "非我行华东三省一市银行汇票");
		docTypeMap.put("A010", "非我行银行本票");
		docTypeMap.put("A011", "特约委托收款");
		docTypeMap.put("A012", "进账单");
		docTypeMap.put("A013", "托收承付(划回)");
		docTypeMap.put("A014", "委托收款(划回)");
		docTypeMap.put("A015", "贷记凭证");
		docTypeMap.put("A016", "特种转账贷方凭证");
		docTypeMap.put("A017", "缴税凭证");
		docTypeMap.put("A018", "财政支付令");
		docTypeMap.put("A019", "代收非税收入凭证");
		docTypeMap.put("A020", "国库资金划缴凭证");
		docTypeMap.put("A021", "收账通知(信用卡付款凭证)");
		docTypeMap.put("A022", "非重空申请书");
		docTypeMap.put("A023", "重空结算申请书");
		docTypeMap.put("A024", "光票托收支票");
		docTypeMap.put("A025", "光票托收旅支");
		docTypeMap.put("A026", "IPCS汇票");
		docTypeMap.put("A027", "国库经收");
	}

	public static String getDocTypeDesc(String docType)
	{
		if (docTypeMap.isEmpty())
		{
			initDocTypeMap();
		}

		try
		{
			return docTypeMap.get(docType) == null ? "" : docTypeMap.get(docType);
		} catch (Exception e)
		{
			return "";
		}
	}

}
