#!/usr/bin/env bash
#nslist=`kubectl  get ns  |grep -v NAME  |awk '{print $1}'`
#for ns in $nslist;do
#        echo $ns
#done
#kubectl label ns default key=fy --overwrite=true