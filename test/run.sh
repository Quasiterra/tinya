#!/bin/bash
failed=0
passed=0
for test in $(find "$1" -name 'test*.sh'); do
    echo "Test $test"
    if ! "$test"; then
        ((failed++))
        echo "FAILED"
    else
        ((passed++))
        echo "PASSED"
    fi
    echo "-----------------------------"
done
echo "$passed PASSED, $failed FAILED"
exit $failed
