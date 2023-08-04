<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户ID:" prop="userid">
          <el-input v-model.number="formData.userid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类型:" prop="typecate">
          <el-input v-model.number="formData.typecate" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in OrderStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="店名:" prop="storename">
          <el-input v-model="formData.storename" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="付款方式:" prop="paymode">
          <el-input v-model="formData.paymode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="金额:" prop="amount">
          <el-input v-model="formData.amount" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="订单号:" prop="orderNo">
          <el-input v-model="formData.orderNo" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:" prop="postscript">
          <el-input v-model="formData.postscript" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="完成时间:" prop="productionedTime">
          <el-date-picker v-model="formData.productionedTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="收件人:" prop="acceptName">
          <el-input v-model="formData.acceptName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机号:" prop="mobile">
          <el-input v-model="formData.mobile" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="街道:" prop="street">
          <el-input v-model="formData.street" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="门牌号:" prop="doorNumber">
          <el-input v-model="formData.doorNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建时间:" prop="createdTime">
          <el-date-picker v-model="formData.createdTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Orders'
}
</script>

<script setup>
import {
  createOrders,
  updateOrders,
  findOrders
} from '@/api/orders'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const OrderStatusOptions = ref([])
const formData = ref({
            userid: 0,
            typecate: 0,
            status: undefined,
            storename: '',
            paymode: '',
            amount: '',
            orderNo: '',
            postscript: '',
            productionedTime: new Date(),
            acceptName: '',
            mobile: '',
            street: '',
            doorNumber: '',
            createdTime: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findOrders({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reorders
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    OrderStatusOptions.value = await getDictFunc('OrderStatus')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createOrders(formData.value)
               break
             case 'update':
               res = await updateOrders(formData.value)
               break
             default:
               res = await createOrders(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
