<template>
    <div class="native_asset_list_container">
        <div class="native_asset_list_header_content">
            <div class="native_asset_list_header_wrap">
                <h2 class="native_header_title_content">
                    Native Assets
                </h2>
            </div>
        </div>
        <div class="native_asset_list_content">
            <div class="native_asset_list_wrap">
                <div>
                    <m-asset-list-table :showNoData="showNoData" :items="nativeAssetList" name="nativeAssetList"></m-asset-list-table>
                </div>
            </div>
            <div v-show="nativeAssetList.length === 0">
                <img class="no_data_img" src="../../assets/no_data.svg">
            </div>
        </div>
    </div>
</template>

<script>
	import Service from "../../service"
	import Tools from "../../util/Tools"
	import MAssetListTable from "./MAssetListTable";
	export default {
		name: "NativeAssetList",
		components: {MAssetListTable},
        data () {
			return {
				showNoData:false,
				nativeAssetList: [],
            }
        },
        mounted(){
		    this.getNativeAssetList()
        },
        methods:{
			getNativeAssetList(){
				Service.commonInterface({nativeAssetList:{}}, (res)=> {
					try {
						if(res){
							this.nativeAssetList = res.data.map( item => {
								return {
									Symbol: item.symbol,
									Owner: item.owner,
									TotalSupply: this.formatNumber(item.total_supply),
									InitialSupply: this.formatNumber(item.initial_supply),
									MaxSupply: this.formatNumber(item.max_supply),
									Mintable: Tools.firstWordUpperCase(item.mintable),
									TokenID: item.token_id,
									flShowLink: true,
                                }
                            })
                        }else {
                        }
					}catch (e) {
                        console.error(e)
					}
                })
            },
            formatNumber(value){
				let million = 1000000;
				if(value > million){
					return `${value/million}M`
                }else {
					return value
                }
            }
        }
	}
</script>

<style scoped lang="scss">
    .native_asset_list_container{
        width: 100%;
        position: relative;
        .no_data_img{
            position: absolute;
            left: 50%;
            top: 50%;
            transform: translate(-50%,-50%);
        }
        .native_asset_list_header_content{
            width: 100%;
            .native_asset_list_header_wrap{
                max-width: 12.8rem;
                margin: 0 auto;
                .native_header_title_content{
                    height: 0.7rem;
                    padding-left: 0.2rem;
                    font-size: 0.18rem;
                    color:#515a6e;
                    line-height: 0.7rem;
                }
            }
        }
        .native_asset_list_content{
            max-width: 12.8rem;
            margin: 0 auto 0.4rem auto;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            .native_asset_list_wrap{
                width: 100%;
                & > div{
                    overflow-x: auto;
                }
            }
        }
    }
    @media screen and (max-width: 910px){
        .native_asset_list_container{
            .native_asset_list_content{
                .native_asset_list_wrap{
                    & > div{
                        margin: 0 0.1rem;
                    }
                }
            }
        }
    }
</style>
