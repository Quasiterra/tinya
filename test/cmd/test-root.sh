#!/bin/bash
res=0
tinya > /dev/null || { res=1; echo "FAIL: should work with no flags!"; }
tinya --badopt &> /dev/null && { res=2; echo "FAIL: should not work with bad option!"; }
exit $res
