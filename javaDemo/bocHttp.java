package com.yizai.banks.boc;

import java.io.UnsupportedEncodingException;

import org.jdom.Document;
import org.jdom.output.Format;

import com.yizai.HandleContext;
import com.yizai.UtilConfigReader;
import com.yizai.banks.BankOperationResult;
import com.yizai.banks.BaseHttpConnParameter;
import com.yizai.banks.HttpBankConnection;
import com.yizai.structure.BaseEnumBankCode;
import com.yizai.structure.BaseEnumEncoding;
import com.yizai.utils.UtilXML;

public class bocHttp
{
	public final static String ENCODING = BaseEnumEncoding.GB18030.getCode();
	public final static String BANKCODE = BaseEnumBankCode.BANKCODE_BOC.getCode();
	private HandleContext hdContext = null;

	public bocHttp(HandleContext hdContext)
	{
		this.hdContext = hdContext;
	}

	public BankOperationResult trade(Document aDoc)
	{
		String reqPacket = null;// 请求报文
		String respPacket = null;// 响应报文
		BankOperationResult objReturn = new BankOperationResult();
		objReturn.setReturnValue(1, "交易成功");

		// 1.加载报文
		if (objReturn.getReturnValue() > 0)
		{
			try
			{
				Format format = Format.getPrettyFormat();
				format.setIndent(" ");
				format.setEncoding(ENCODING);
				format.setExpandEmptyElements(false);
				reqPacket = UtilXML.docToString(aDoc, format);
			} catch (Exception e)
			{
				
				hdContext.BankLogError(e.getMessage(), e);
				objReturn.setReturnValue(-1, hdContext.getResValue("msg.bank_get_send_msg_fail"));
			}
		}
		// 2.发送请求
		if (objReturn.getReturnValue() > 0)
		{
			try
			{
				hdContext.BankLogInfo(hdContext.getResValue("msg.bank_send_msg"));
				hdContext.BankLogInfo(reqPacket);
				// 初始化HTTP请求参数
				BaseHttpConnParameter connPara = initHttpConnParam(reqPacket);
				// 发送交易报文
				respPacket = send(reqPacket, connPara, hdContext);
				// System.out.println("响应报文：\n" + respPacket);
			} catch (Exception e)
			{
				
				hdContext.BankLogError(e.getMessage(), e);
				objReturn.setReturnValue(-1, hdContext.getResValue("msg.bank_get_send_msg_fail"));
			}
		}
		// 3.解析银行返回报文
		if (objReturn.getReturnValue() > 0)
		{
			try
			{
				hdContext.BankLogInfo(hdContext.getResValue("msg.bank_receive_msg"));
				hdContext.BankLogInfo(respPacket);
				// 组装响应报文
				Document responsePacket = UtilXML.stringToDoc(respPacket, ENCODING);
				objReturn.setResponseDoc(responsePacket);
			} catch (Exception e)
			{
				
				hdContext.BankLogError(e.getMessage(), e);
				objReturn.setReturnValue(-1, hdContext.getResValue("msg.bank_get_send_msg_fail"));
			}
		}
		return objReturn;
	}

	private static BaseHttpConnParameter initHttpConnParam(String reqPacket) throws UnsupportedEncodingException
	{
		BaseHttpConnParameter connParameter = new BaseHttpConnParameter();
		String ip = UtilConfigReader.getBcpConfig(BANKCODE, "ip");
		String port = UtilConfigReader.getBcpConfig(BANKCODE, "port");
		connParameter.setIp(ip);
		connParameter.setPort(port);
		connParameter.setUri("/B2EC/E2BServlet");
		connParameter.setMethod("POST");
		connParameter.setAccept("image/gif, image/x-xbitmap, image/jpeg, image/pjpeg, application/vnd.ms-powerpoint, application/vnd.ms-excel, application/msword, */*");
		connParameter.setAcceptLanguage("zh-cn");
		connParameter.setContentType("text/plain");
		connParameter.setProxyConnection("Keep-Alive");
		connParameter.setUserAgent("Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0)");
		connParameter.setHost(" " + connParameter.getIp() + ":" + connParameter.getPort() + "");
		connParameter.setContentLength("" + Integer.toString(reqPacket.getBytes(ENCODING).length, 10));
		connParameter.setPragma("no-cache");
		return connParameter;
	}

	private static String send(String aSendData, BaseHttpConnParameter aConnRow, HandleContext hdContext) throws Exception
	{
		String strRtn = "";
		HttpBankConnection conn = null;
		try
		{
			conn = new HttpBankConnection(hdContext);// 初始化连接
			conn.connect(aConnRow);// 连接
			conn.send(aSendData, ENCODING);// 发送
			strRtn = conn.receive(ENCODING);// 接收返回
		} finally
		{
			if (conn != null)
			{
				conn.disConnect();// 断开连接
			}
		}
		return strRtn;
	}
}
