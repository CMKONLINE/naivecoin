package controllers

import (
	"encoding/json"
	"naivecoin/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about Blockchain
type BlockchainController struct {
	beego.Controller
}

// @Title Get
// @Description find Blockchain by index
// @Param	index	path 	int 	true		"the index you want to get"
// @Success 200 {Block} models.Block
// @Failure 403 :index is empty
// @router /:index [get]
func (o *BlockchainController) Get() {
	index := o.Ctx.Input.Param(":index")
	indexBlock, err := strconv.Atoi(index)
	if err == nil {
		ob := models.GetBlockByIndex(indexBlock)
		o.Data["json"] = ob
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all blocks
// @Success 200 {BlockChain} models.BlockChain
// @Failure 403 :BlockchainId is empty
// @router / [get]
func (o *BlockchainController) GetAll() {
	obs := models.GetAllBlocks()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the latest block
// @Param	body		body 	models.Block	true		"The body"
// @Success 200 {Block} models.Block
// @Failure 403 :Block is empty
// @router /:BlockchainId [put]
func (o *BlockchainController) Put() {
	var block models.Block
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &block)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		if models.AddBlock(block) {
			o.Data["json"] = "update success!"
		}
	}
	o.ServeJSON()
}
