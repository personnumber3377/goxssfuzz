#!/bin/sh

#./fuzzer -max_len=100 -only_ascii=1 -dict=dict.dict corp/


 ./fuzzer -max_len=100 -only_ascii=1 -timeout=20 -dict=dict.dict corp/

#while :
#do
#	# ./fuzzer -max_len=100 -only_ascii=1 -fork=1 -ignore_crashes=1 -ignore_timeouts=1 -ignore_ooms=1 -timeout=20 -dict=dict.dict corp/ 2>> fuzz_output.txt || true # We need this because the net/html is a fucking buggy mess which has many infinite loops.
#done

