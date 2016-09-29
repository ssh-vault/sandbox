Convert RSA keys to PEM:

    $ openssl rsa -in ~/.ssh/id_rsa -pubout -outform pem > id_rsa.pub.pem

Generate a password, encrypt the file with it symmetrically, and encrypt the password with your public, key saving it to file:

    $ openssl rand 64 | tee >(openssl enc -aes-256-cbc -pass stdin -in file.txt -out file.enc) | openssl rsautl -encrypt -pubin -inkey id_rsa.pub.pem  -out file.enc.key


Decrypt the passphrase with your private key and use it to decrypt the file:

    $ openssl rsautl -decrypt -inkey ~/.ssh/id_rsa -in file.enc.key | openssl enc -aes-256-cbc -pass stdin -d -in file.enc -out file.txt
