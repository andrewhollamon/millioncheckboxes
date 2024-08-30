# SCALING CONSIDERATIONS

This is a list of other steps we could take, to potentially increase throughput or performance, 
at the various layers.
- uuid pre-generation pool server 
- db: separate storage pools for each partition
- db: separate storage pools for the hot table vs the rest, or also for the updates table 
- 