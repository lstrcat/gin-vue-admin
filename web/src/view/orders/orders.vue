<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
            <el-switch
              v-model="swAudio"
              @change="handleOpenAudio"
              class="ml-2"
              inline-prompt
              style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
              active-text="开启声音"
              inactive-text="音频"
            />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        :row-class-name="tableRowClassName"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="用户ID" prop="userid" width="80" />
        <el-table-column align="left" label="收件人" width="100">
          <template #default="scope">
          <el-popover effect="light" trigger="hover" placement="top" width="auto">
            <template #default>
              <div>{{scope.row.mobile}}</div>
              <div>{{scope.row.street }}</div>
              <div>{{scope.row.doorNumber }}</div>
            </template>
            <template #reference>
              <el-tag>{{ scope.row.acceptName}}</el-tag>
            </template>
          </el-popover>
        </template>
        </el-table-column>
        <el-table-column align="left" label="状态" width="100">
          <template #default="scope">
            <template v-if="scope.row.status == 1">
              <el-tag class="ml-2" type="success" >已下单</el-tag>
            </template>             
            <template v-else-if="scope.row.status == 2">
              <el-tag class="ml-2" type="info" >装车中</el-tag>
            </template>            
            <template v-else-if="scope.row.status == 3">
              <el-tag class="ml-2" type="warning" >配送中</el-tag>
            </template>            
            <template v-else-if="scope.row.status == 4">
              <el-tag class="ml-2" type="danger" >已送达</el-tag>
            </template>
           </template>
        </el-table-column>
        <el-table-column align="left" label="金额" prop="amount" width="120" />
        <el-table-column align="left" label="订单号" prop="orderNo" width="150" />
         <el-table-column align="left" label="完成时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.productionedTime) }}</template>
         </el-table-column>
        <el-table-column width="200" label="操作">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateOrdersFunc(scope.row)">详情</el-button>
            <!--<el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>-->
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'打印订单':'打印订单'" destroy-on-close>
      <div id="printContent">
      <el-descriptions
          :model="formData"
          direction="horizontal"
          :column="2"
          :size="size"
          border>
            <el-descriptions-item label="订单编号">{{formData.orderNo}}</el-descriptions-item>
            <el-descriptions-item label="创建时间">{{formatDate(formData.CreatedAt)}}</el-descriptions-item>
            <el-descriptions-item label="用户ID">{{formData.userid}}</el-descriptions-item>
            <el-descriptions-item label="付款方式" :span="2">{{formData.paymode}}</el-descriptions-item>
            <el-descriptions-item label="金额">{{formData.amount}}</el-descriptions-item>
            <el-descriptions-item label="收件人">
              {{formData.acceptName + ' ' + formData.mobile}}
            </el-descriptions-item>
            <el-descriptions-item label="地址"
              >{{formData.street + ' ' + formData.doorNumber}}
            </el-descriptions-item>
            <el-descriptions-item label="备注">
              <el-tag class="ml-2" type="danger" >{{formData.postscript}}</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="状态">
              <el-select v-model="formData.status" placeholder="请选择" >
                <el-option v-for="(item,key) in OrderStatusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-descriptions-item>
      </el-descriptions>    
      <el-table 
        :data="tableDataGoods" 
        row-key="ID"
        border="true"
        style="width: 100%">
          <el-table-column prop="name" label="品名" width="180" />
          <el-table-column prop="property" label="属性" width="auto" />    
          <el-table-column prop="price" label="价格" width="80" />   
          <el-table-column prop="number" label="数量" width="80" />                  
      </el-table>    
      </div>    
      <template #footer>
        <div class="dialog-footer">
          <el-button v-print="printObj" type="primary" >打印</el-button>
          <el-button @click="closeDialog">取 消</el-button>
          <!--
          <el-button type="primary" @click="enterDialog">确 定</el-button>-->
        </div>
      </template>
    </el-dialog>
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
  deleteOrders,
  deleteOrdersByIds,
  updateOrders,
  findOrders,
  getOrdersList
} from '@/api/orders'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox, ElNotification  } from 'element-plus'
import { ref, reactive, onMounted, onDeactivated } from 'vue'

// 自动化生成的字典（可能为空）以及字段
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
const tableDataGoods = ref([])
const swAudio = ref(false)

// 打印对象
const printObj = {
  id: "printContent",
  popTitle: "沙场送货单",
};

// 订单通知
const orderNotification = (title,str) => {
  ElNotification({
    title: title,
    message: str,
    duration: 0,
  })
  let audio = new Audio()
	audio.src = "http://ryhsxbzuv.bkt.clouddn.com/mp3/hll.mp3"
  audio.play()
}
			

// 定时器
const timer = ref(0)
const index = ref(0)
onMounted(()=>{ //组件挂载时的生命周期执行的方法
    timer.value = window.setInterval(function logname() {
        // 其他定时执行的方法
        index.value = index.value + 1
        console.log("刷新订单列表"+ index.value);
        checkGetTableData()
    }, 5000);
})
  onDeactivated(()=>{ //离开当前组件的生命周期执行的方法
    window.clearInterval(timer.value);
})

// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 表格行颜色
const tableRowClassName = ({
  row,
  rowIndex,
}) => {
  if (rowIndex === 0 && row.newcomer == true) {
    return 'warning-row'
  } else if (rowIndex === 3) {
    return ''
  }
  return ''
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询-是否有新订单
const checkGetTableData = async() => {
  const table = await getOrdersList({ page: 1, pageSize: 10 })

  if (table.code === 0) {
    let newOrder = {}
    if(table.data.list.length > 0){
      newOrder = table.data.list[0]
    }
    let oldOrder = {}
    if(tableData.value.length > 0){
      oldOrder = tableData.value[0]
    }
    if( newOrder.ID != oldOrder.ID){
      console.log("新订单: ",newOrder.ID, " 旧订单: ", oldOrder.ID)
      newOrder.newcomer = true
      orderNotification('新订单', '您有新的订单，请注意查收' + '(' + newOrder.orderNo + ')')

      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }

  }
}

// 查询
const getTableData = async() => {
  const table = await getOrdersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    OrderStatusOptions.value = await getDictFunc('OrderStatus')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

const handleOpenAudio = () => {
  if(swAudio.value == true) {
   orderNotification("提示","打开新订单声音通知")
  }
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteOrdersFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteOrdersByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateOrdersFunc = async(row) => {
    const res = await findOrders({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reorders
        tableDataGoods.value = res.data.goods
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteOrdersFunc = async (row) => {
    const res = await deleteOrders({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>
.el-table .warning-row {
  --el-table-tr-bg-color: var(--el-color-warning-light-9);
}
.el-table .success-row {
  --el-table-tr-bg-color: var(--el-color-success-light-9);
}
</style>
