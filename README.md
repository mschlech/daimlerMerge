# Daimler Merge

Implementation of a MERGE Function which takes a list of intervals which returns a list of all non overlapping intervalls subsequently.

example:
Input: [25,30] [2,19] [14, 23] [4,8]  Output: [2,23] [25,30]

Runtime of the implementation ? 
How can a robustness of the implmentation being fullfilled of huge amount of incoming lists ?
Memory consumption ?



![pipeline status](https://gitlab.com/mschlechdaimer_merge/badges/master/pipeline.svg)](https://gitlab.com/mschlech/daimler_merge)

# start the docker container 
CONTAINER_NAME = docker images
docker run -t -i -p 3000:3000 -ti --rm --init <containerName>


## run
within your IDE run the Main Method 

http://localhost:8080/mergedList

## Time complexity and resulting computational drawbacks

The interval to be found is the result of comparing 2 values. The first one of a tuple and the second one. The first one is the offset in the code and the second one the end. The notation may be misleading.
 One complexity is the O(n2) which can be optimized by sorting which results in a O(logn) situation. 

![alg](doc/images/Alg.png)
