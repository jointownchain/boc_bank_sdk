package com.yizai.banks.boc;

import java.text.SimpleDateFormat;
import java.util.Date;

import org.apache.commons.lang.StringUtils;
import org.jdom.Document;

import com.yizai.HandleBankUtility;
import com.yizai.HandleContext;
import com.yizai.HandleResult;
import com.yizai.HandleUtility;
import com.yizai.Req_Pain_001_001_03_Input;
import com.yizai.Req_Pain_001_001_03_Output;
import com.yizai.banks.BankOperationResult;
import com.yizai.structure.BaseEnumBankCode;
import com.yizai.structure.BaseEnumChrgBr;
import com.yizai.structure.BaseEnumIsoCurrency;
import com.yizai.structure.BaseEnumTxSts;
import com.yizai.utils.DateUtils;
import com.yizai.utils.UtilXML;

public class bocPain00100103Foreign
{
	//外币付款
	public static void ReqPain00100103(Req_Pain_001_001_03_Input input, Req_Pain_001_001_03_Output output, HandleResult handleResult, HandleContext contextInput, BaseEnumBankCode bankCode)
	{
		SimpleDateFormat dataFormater = new SimpleDateFormat(DateUtils.DATETIME_FORMAT_DATE_WITH_DASH);

		//Step 0, 银行的额外验证
		if ("CN".equals(input.getCdtrAgt().getCountry()) && (input.getCcyOfTrf().equals(BaseEnumIsoCurrency.RMB) || input.getCcyOfTrf().equals(BaseEnumIsoCurrency.CNY)))
		{
			//仅支持外币
			handleResult.setFailed(contextInput.getResValue("msg.interface_only_support_foreign_curr"));
			return;
		}

		//Step 1, 加载XML报文
		String template = "b2e0200";
		Document requestXMLDoc = HandleUtility.getResponseXML(template, bocImpl.requestXMLs, contextInput, handleResult);

		//Step 2, 生成发送报文
		if (handleResult.isSucc())
		{
			try
			{
				// 初始化报文头
				bocUtil.setHeadForRequest(requestXMLDoc, bankCode);
				// 获取文档根
				String transPath = "";
				transPath = "/bocb2e/trans/trn-" + template + "-rq/" + template + "-rq";
				String auth_date = dataFormater.format(new Date());
				UtilXML.setText(requestXMLDoc, transPath + "/AUTH_EXPIRY_DT", auth_date);//授权有效日期 
				UtilXML.setText(requestXMLDoc, transPath + "/REM_MODE", "S");//汇款类型 
				UtilXML.setText(requestXMLDoc, transPath + "/RET_LVL", "1");//发电等级
				UtilXML.setText(requestXMLDoc, transPath + "/TRN_AMT", input.getInstdAmt().toPlainString());//汇款金额
				//中行的币种需要传入国际标准的英文大写字母
				UtilXML.setText(requestXMLDoc, transPath + "/TRN_CUR", input.getCcyOfTrf().toString());//汇款币种

				if (input.getCcyOfTrf().equals(input.getDbtrAcct().getCcy()))
				{
					UtilXML.setText(requestXMLDoc, transPath + "/CEX_AMT", input.getInstdAmt().toPlainString());//现汇金额
					UtilXML.setText(requestXMLDoc, transPath + "/CEX_ACNO", input.getDbtrAcct().getId());//现汇帐号					
				} else
				{
					UtilXML.setText(requestXMLDoc, transPath + "/PUR_AMT", input.getInstdAmt().toPlainString());//购汇金额
					UtilXML.setText(requestXMLDoc, transPath + "/PUR_ACNO", input.getDbtrAcct().getId());//购汇帐号
				}
				UtilXML.setText(requestXMLDoc, transPath + "/FRN_NAME_50A", input.getDbtrAcct().getNm());//50a汇款人名址
				String Rmt_Psn_Adr = input.getDbtr().getAdrLine() + "," + input.getDbtr().getTwnNm() + "," + input.getDbtr().getCtrySubDvsn() + "," + input.getDbtr().getCtry();
				UtilXML.setText(requestXMLDoc, transPath + "/FRN_ADDRESS_50A", Rmt_Psn_Adr);//50a汇款人地址

				String organCode = input.getOrganCode();
				if (organCode == null || "".equals(organCode))
				{
					handleResult.setFailed(contextInput.getResValue("msg.organ_required_for_payment", input.getEndToEndId()));
					return;
				} else
				{
					UtilXML.setText(requestXMLDoc, transPath + "/ORG_CDE", organCode);//组织机构代码	
				}

				String creditBankName = input.getCdtrAgt().getNm();
				if (creditBankName == null || "".equals(creditBankName))
				{
					handleResult.setFailed(contextInput.getResValue("msg.payee_bank_name_required_for_payment", input.getEndToEndId()));
					return;
				} else
				{
					UtilXML.setText(requestXMLDoc, transPath + "/BEAK_NAME", creditBankName);//57a收款银行名址	
				}

				if (input.getCdtrAcct().getIBAN() != null && !"".equals(input.getCdtrAcct().getIBAN()))
				{
					UtilXML.setText(requestXMLDoc, transPath + "/TO_ACT_ACN_59A", input.getCdtrAcct().getIBAN());//收款人IBAN帐号	
				} else
				{
					UtilXML.setText(requestXMLDoc, transPath + "/TO_ACT_ACN_59A", input.getCdtrAcct().getId());//收款人帐号					
				}

				UtilXML.setText(requestXMLDoc, transPath + "/TO_NAME_59A", input.getCdtrAcct().getNm());//收款人名称
				String RcvPymtPs_Adr = input.getCdtr().getAdrLine() + "'," + input.getCdtr().getTwnNm() + "," + input.getCdtr().getCtrySubDvsn() + "," + input.getCdtr().getCtry();
				UtilXML.setText(requestXMLDoc, transPath + "/TO_ADDRESS_59A", RcvPymtPs_Adr);//收款人地址

				String usage = HandleBankUtility.getPurposeStrWithEng(input.getInstrPrty(), input.getPurpCd(), input.getEndToEndId());
				UtilXML.setText(requestXMLDoc, transPath + "/POSTSCRIPTMT", usage);//汇款附言

				String chargerBear = "";
				if (input.getChrgBr().equals(BaseEnumChrgBr.CRED))
				{
					chargerBear = "BEN";//收款人承担
				} else if (input.getChrgBr().equals(BaseEnumChrgBr.DEBT))
				{
					chargerBear = "OUR";//付款人承担
				} else if (input.getChrgBr().equals(BaseEnumChrgBr.SHAR))
				{
					chargerBear = "SHA";//共同承担
				} else
				{
					handleResult.setFailed(contextInput.getResValue("msg.charger_bear_required", input.getEndToEndId()));
					return;
				}
				UtilXML.setText(requestXMLDoc, transPath + "/FEE_MODE", chargerBear);//71a国内外费用承担

				String countryCode = "";
				countryCode = HandleBankUtility.get3DigiCountryCode(input.getCdtr().getCtry());
				if (countryCode == null || "".equals(countryCode))
				{
					//中国银行最新的版本中，该字段可以为空
					//					handleResult.setFailed(contextInput.getResValue("boc.country_code_read_failed", input.getEndToEndId(), countryCode));
					//					return;
				} else
				{
					UtilXML.setText(requestXMLDoc, transPath + "/COU_CDE", countryCode);//收款人常驻国 家(地区)代码
				}

				//申请人名称
				String pername = input.getBopCode().getApplicantName().trim();
				if (pername == null || "".equals(pername))
				{
					handleResult.setFailed(contextInput.getResValue("msg.application_name_required_for_payment", input.getEndToEndId(), countryCode));
					return;
				} else
				{
					UtilXML.setText(requestXMLDoc, transPath + "/PER_NAME", pername);//申请人名称	
				}

				//申请人联系电话
				String perphone = input.getBopCode().getApplicantPhone().trim();
				if (perphone == null || "".equals(perphone))
				{
					handleResult.setFailed(contextInput.getResValue("msg.application_phone_required_for_payment", input.getEndToEndId(), countryCode));
					return;
				} else
				{
					UtilXML.setText(requestXMLDoc, transPath + "/PER_TEL", perphone);//申请人联系电话
				}

				UtilXML.setText(requestXMLDoc, transPath + "/ERP_REF", input.getEndToEndId());//集团业务编号

				//<REM_TYP> 汇款种类
				//非空 1-国外汇 款；2-国内汇 款
				String transType = "";
				if (input.getCdtrAgt().getCountry().equals("CN"))
				{
					transType = "2";
				} else
				{
					transType = "1";
				}
				UtilXML.setText(requestXMLDoc, transPath + "/REM_TYP", transType);//汇款种类
				UtilXML.setText(requestXMLDoc, transPath + "/IBK_NUM", bocUtil.setCityCodeForBOC(input.getDbtrAgt().getMmbId()));//联行号
				UtilXML.setText(requestXMLDoc, transPath + "/CUS_FLAG", "BN");//客户标识

				if (input.getCdtrAgt().getBIC() != null && !"".equals(input.getCdtrAgt().getBIC()))
				{
					UtilXML.setText(requestXMLDoc, transPath + "/BED_BIC_57A", input.getCdtrAgt().getBIC());//<BED_BIC_57A> 57a收款银行 BIC
				}

				//BOP处理
				// 处理BOP Code
				if (!HandleBankUtility.isBopEmpty(input.getBopCode()))
				{
					//选输 1-预付货款 2-货到付款 3-退款 4-其他
					//CITI的Payment Purpose(1位标志 A–Advance Payment D–Payment Against Delivery R–Refund O–Others)
					String PAY_TYPE = "";
					if (input.getBopCode().getPayPurpose().equals("A"))
					{
						PAY_TYPE = "1";
					} else if (input.getBopCode().getPayPurpose().equals("D"))
					{
						PAY_TYPE = "2";
					} else if (input.getBopCode().getPayPurpose().equals("R"))
					{
						PAY_TYPE = "3";
					} else if (input.getBopCode().getPayPurpose().equals("O"))
					{
						PAY_TYPE = "4";
					} else
					{
						handleResult.setFailed(contextInput.getResValue("msg.bop_pay_purpose_wrong_for_payment", input.getBopCode().getPayPurpose(), input.getEndToEndId()));
						return;
					}
					UtilXML.setText(requestXMLDoc, transPath + "/PAY_TYPE", PAY_TYPE);
					UtilXML.setText(requestXMLDoc, transPath + "/TRN_CD1", input.getBopCode().getTransCode());//交易编码1
					UtilXML.setText(requestXMLDoc, transPath + "/TRN_CUR1", input.getCcyOfTrf().toString());//交易币别
					UtilXML.setText(requestXMLDoc, transPath + "/TRN_AMT1", input.getInstdAmt().toPlainString());//交易金额1
					UtilXML.setText(requestXMLDoc, transPath + "/TRN_RE1", input.getBopCode().getTransCodeRemark());//交易附言1

					//是否保税区货物
					String BGD_PAY = "";
					if (input.getBopCode().getBondedGoods().equals("Y"))
					{
						BGD_PAY = "1";
					} else if (input.getBopCode().getBondedGoods().equals("N"))
					{
						BGD_PAY = "0";
					} else
					{
						handleResult.setFailed(contextInput.getResValue("msg.bop_bonded_goods_wrong_for_payment", input.getBopCode().getBondedGoods(), input.getEndToEndId()));
						return;
					}
					UtilXML.setText(requestXMLDoc, transPath + "/IS_BGD_PAY", BGD_PAY);//本币款项是否 为保税货物项下付款
					UtilXML.setText(requestXMLDoc, transPath + "/COM_NUM", input.getBopCode().getContractNumber());//合同号
					UtilXML.setText(requestXMLDoc, transPath + "/INV_NUM", input.getBopCode().getInvoiceNumber());//发票号
				}
			} catch (Exception e)
			{
				handleResult.setFailed(contextInput.getResValue("msg.build_send_message_failed", e.getMessage(), "", "", ""));
				contextInput.BankLogError(e.getMessage(), e);
				return;
			}
		}

		//Step 3 发送银行
		Document returnDocument = null;
		if (handleResult.isSucc())
		{
			BankOperationResult opResult = new bocHttp(contextInput).trade(requestXMLDoc);
			if (opResult.getReturnValue() > 0)
			{
				returnDocument = (Document) opResult.getResponseDoc();
				if (returnDocument == null)
				{
					//此时已经有可能发送成功了，不允许再设置成ERROR
					output.setTxSts(BaseEnumTxSts.PDNG);
					String message = contextInput.getResValue("msg.build_retrun_message_failed");
					output.setAddtlInf(contextInput.getResValue(message));
					contextInput.BankLogError(message);
					return;
				}
			} else
			{
				//此时已经有可能发送成功了，不允许再设置成ERROR
				output.setTxSts(BaseEnumTxSts.PDNG);
				String message = contextInput.getResValue("msg.build_return_message_failed_due_to", opResult.getErrorDetailMessage());
				output.setAddtlInf(contextInput.getResValue(message));
				contextInput.BankLogError(message);
				return;
			}
		}

		//Step 4 处理结果
		if (handleResult.isSucc())
		{
			try
			{
				// 获取交易代码
				String strTranCode = UtilXML.queryChildText(returnDocument.getRootElement(), "./head/trncod");
				// 如果返回的交易代码为“严重错误”返回报文
				if ("b2eerror".equals(strTranCode) || StringUtils.isBlank(strTranCode))
				{
					String transPath = "/" + returnDocument.getRootElement().getName() + "/trans/trn-b2eerror-rs";
					String respMsg = UtilXML.queryChildText(returnDocument, transPath + "/status/rspmsg");
					output.setTxSts(BaseEnumTxSts.PDNG);
					output.setAddtlInf(respMsg);
				} else
				{
					String transPath = "/" + returnDocument.getRootElement().getName() + "/trans/trn-" + template + "-rs";
					String respCode = UtilXML.queryChildText(returnDocument, transPath + "/status/rspcod");
					String respMsg = UtilXML.queryChildText(returnDocument, transPath + "/status/rspmsg");
					if (respCode.equals("B001"))
					{
						respCode = UtilXML.queryChildText(returnDocument, transPath + "/b2e0200-rs/status/rspcod");// 交易流水状态码
						respMsg = UtilXML.queryChildText(returnDocument, transPath + "/b2e0200-rs/status/rspmsg"); // 交易流水状态消息

						// 首先判断被交易流水状态码是否为空,不为空再对其状态进行判断;否则置为不明,返回报文处理的返回码和返回消息
						if (respCode != null && !respCode.trim().equals(""))
						{
							if (respCode != null && respCode.equals("B001"))
							{
								String OBSS_ID = UtilXML.queryChildText(returnDocument, transPath + "/b2e0200-rs/OBSS_ID");
								output.setTxSts(BaseEnumTxSts.PDNG);
								output.setAddtlInf(respMsg);
								output.setOrgnlInstrId(OBSS_ID);
								output.setOrgnlEndToEndId(input.getEndToEndId());
							} else
							{
								output.setTxSts(BaseEnumTxSts.PDNG);
								output.setAddtlInf(respMsg);
							}
						} else
						{
							output.setTxSts(BaseEnumTxSts.PDNG);
							output.setAddtlInf(respMsg);
						}
					} else
					{
						output.setTxSts(BaseEnumTxSts.PDNG);
						output.setAddtlInf(respMsg);
					}
				}
			} catch (Exception ex)
			{
				output.setTxSts(BaseEnumTxSts.PDNG);
				output.setAddtlInf(contextInput.getResValue("msg.build_return_message_failed_due_to", ex.getMessage()));
				contextInput.BankLogError(ex.getMessage(), ex);
			}
		}
	}
}
