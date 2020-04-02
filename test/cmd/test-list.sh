#!/bin/bash
res=0

tinya list > /dev/null || { res=1; echo "FAIL: list"; }

AWS_ACCESS_KEY_ID=who AWS_SECRET_ACCESS_KEY=me \
    tinya list &> /dev/null && { res=2; echo "FAIL: bad list"; }

exit $res
