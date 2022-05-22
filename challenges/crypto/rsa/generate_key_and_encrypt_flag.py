import base64
from Crypto.Cipher import PKCS1_OAEP
from Crypto.PublicKey import RSA


FLAG = "FLAG{W1k1ped1a_1s_y0ur_Fr1enD!}"

# https://zh.wikipedia.org/zh-tw/RSA%E5%8A%A0%E5%AF%86%E6%BC%94%E7%AE%97%E6%B3%95
n = 1230186684530117755130494958384962720772853569595334792197322452151726400507263657518745202199786469389956474942774063845925192557326303453731548268507917026122142913461670429214311602221240479274737794080665351419597459856902143413
e = 65537

public_key = RSA.construct((n, e))

with open("src/key.pub", "wb+") as f:
    f.write(public_key.export_key("PEM"))

cipher_rsa = PKCS1_OAEP.new(public_key)

encrypt_flag = cipher_rsa.encrypt(FLAG.encode())
encrypt_flag = base64.b64encode(encrypt_flag).decode()

with open("src/flag.enc", "w+") as f:
    f.write(encrypt_flag)
