### DFS: Distributed File System

2018/06/04  Chao
This is a simplified version of ipfs. The intend is to understand and learn how ipfs work. GPLv2 License. 

example:

1. add a file:  ./dfs add map.jpg  (will give hash QmZUyJE41WoEMKUNWzv97LvE4YbSEE63wGp2peneq1ojf7)
2. retrive a file from the hash: ./dfs cat QmZUyJE41WoEMKUNWzv97LvE4YbSEE63wGp2peneq1ojf7 > map1.jpg
3. check they are equal: diff map.jpg map1.jpg


