#!/usr/bin/env bash

host="http://jungle2.cryptolions.io"
contractAcc="claim1111112"

#remove code: cleos set contract <account> --clear
echo "host: " 	  $host
echo "contract: " $contractAcc

printEvent(){
	echo "sale market:"
	cleos -u $host get table $contractAcc $contractAcc event
	echo "account info:"
	cleos -u $host get account $contractAcc
}

case "$1" in
"deploy")
	echo "----- deploy"
	cd ..
	sh build.sh
	echo "account info:"
	cleos -u $host get account $contractAcc
	echo "deploy:"
	cleos -u $host set contract $contractAcc $(pwd)/build claims.wasm claims.abi
	echo "check code:"
	cleos -u $host get code $contractAcc
	echo "account info:"
	cleos -u $host get account $contractAcc
    ;;
"hi")
	echo "----- action hi"
	cleos -u $host push action $contractAcc hi '["myacc1111111"]' -p $contractAcc@active
    printEvent
	;;

esac


