#!/usr/bin/env bash
nslist=`kubectl  get ns  |grep -v NAME  |awk '{print $1}'`
for ns in $nslist;do
        kubectl label  microservices.asm.alauda.io --all  app.cpaas.io/microservice-type=service-mesh  -n $ns
        list=`kubectl  get microservices.asm.alauda.io -n $ns  |grep -v NAME  |awk '{print $1}'`
        for ms in $list;do
                echo $ms $ns
                kubectl -n $ns patch microservices.asm.alauda.io $ms --type='json' -p='[{ "op": "replace", "path": "/apiVersion", "value": "asm.alauda.io/v1beta3" }]'
        done
done