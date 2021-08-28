# mcoin
golang mcoin 实现拉取矿币信息

# use
具体用法:
mcoin 
  -k string
        <btc/eth/etc>
# view
[root@chengkenlee ~]# mcoin -k btc

"币种标识","英文名称","中文名称","交易对","当前价格","被兑换货币交易量","兑换货币交易量","货币类型","涨跌百分比","24小时趋势","货币供应量","总市值","趋势数据" 

SYMBOL  NAME_EN   PAIR      RATE           VOL_A             VOL_B        CURR_SUFFIX  RATE_PERCENT TREND     SUPPLY         MARKETCAP         LQ   NAME_CN 
BTC     Bitcoin   btc_usdt  48873.74       534.17350176276   25,950,357    USDT        3.21         up        18797756       918,716,639,327   0    比特币 

[root@chengkenlee ~]#
![image](https://github.com/chengkenlee/img/blob/master/%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20210828212811.png)





