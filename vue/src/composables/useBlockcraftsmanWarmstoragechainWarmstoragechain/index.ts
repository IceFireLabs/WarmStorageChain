/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/vue-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useBlockcraftsmanWarmstoragechainWarmstoragechain() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.BlockcraftsmanWarmstoragechainWarmstoragechain.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryResource = (index: string,  options: any) => {
    const key = { type: 'QueryResource',  index };    
    return useQuery([key], () => {
      const { index } = key
      return  client.BlockcraftsmanWarmstoragechainWarmstoragechain.query.queryResource(index).then( res => res.data );
    }, options);
  }
  
  const QueryResourceAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryResourceAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.BlockcraftsmanWarmstoragechainWarmstoragechain.query.queryResourceAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryParams,QueryResource,QueryResourceAll,
  }
}