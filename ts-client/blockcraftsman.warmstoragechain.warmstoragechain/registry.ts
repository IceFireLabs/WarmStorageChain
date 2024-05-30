import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateResource } from "./types/warmstoragechain/warmstoragechain/tx";
import { MsgCreateResource } from "./types/warmstoragechain/warmstoragechain/tx";
import { MsgDeleteResource } from "./types/warmstoragechain/warmstoragechain/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/blockcraftsman.warmstoragechain.warmstoragechain.MsgUpdateResource", MsgUpdateResource],
    ["/blockcraftsman.warmstoragechain.warmstoragechain.MsgCreateResource", MsgCreateResource],
    ["/blockcraftsman.warmstoragechain.warmstoragechain.MsgDeleteResource", MsgDeleteResource],
    
];

export { msgTypes }