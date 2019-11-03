> perf stat -e cycles,instructions,page-faults,cache-misses,branches,branch-misses ./vec

 Performance counter stats for './vec':

       281,665,334      cycles                                                      
     1,026,730,077      instructions              #    3.65  insn per cycle         
             2,543      page-faults                                                 
            80,625      cache-misses                                                
       235,035,846      branches                                                    
           213,968      branch-misses             #    0.09% of all branches        

       0.110326499 seconds time elapsed
       
       
> perf stat -e cycles,instructions,page-faults,cache-misses,branches,branch-misses ./rowbased 

 Performance counter stats for './rowbased':

     2,586,719,384      cycles                                                      
     7,385,760,814      instructions              #    2.86  insn per cycle         
             4,998      page-faults                                                 
         1,037,583      cache-misses                                                
     1,441,232,904      branches                                                    
         1,608,828      branch-misses             #    0.11% of all branches        

       0.871683493 seconds time elapsed
       

从上面的对比可以看出：
1. 向量化执行更快: 耗时是非向量化的 1/8 左右；
2. 向量化执行用更少的指令完成了相同的工作: 因为降低了解释开销，减少了"无意义"的指令数量；
3. 向量化执行对内存更加友好: 因为大多数时候是在集中的访问连续内存，所以能更好的命中 CPU Cache；
4. 向量化执行对 CPU 分支预测更有利；