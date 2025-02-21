<template>
    <div class="order-management">
        <el-card class="toolbar-card">
            <div class="toolbar">
                <div class="filters">
                    <!-- 新增商户选择器 -->
                    <el-select v-model="searchParams.merchantId" :placeholder="t('common.selectMerchant')" clearable
                        filterable style="width: 200px; margin-right: 15px" @change="loadData">
                        <el-option v-for="merchant in merchantOptions" :key="merchant.value" :label="merchant.label"
                            :value="merchant.value" />
                    </el-select>
                    <el-input v-model="searchParams.keyword" :placeholder="t('order.searchOrder')" style="width: 300px"
                        clearable @clear="loadData">
                        <template #append>
                            <el-button icon="el-icon-search" @click="loadData" />
                        </template>
                    </el-input>
                </div>
            </div>
        </el-card>
        <el-card class="table-card" style="margin-top: 20px">
            <el-table :data="orderList" border stripe v-loading="loading" style="width: 100%" :max-height="500">
                <el-table-column prop="no" :label="t('order.orderNo')" :min-width="160" show-overflow-tooltip />
                <el-table-column :label="t('order.merchantName')" :min-width="150" show-overflow-tooltip>
                    <template #default="{ row }">{{ getMerchantName(row.m_id) }}</template>
                </el-table-column>
                <el-table-column prop="amount" :label="t('order.amount')" align="right" :width="140">
                    <template #default="{ row }">{{ row.amount }} {{ row.coin }}</template>
                </el-table-column>
                <el-table-column prop="status" :label="t('common.status')" :width="160">
                    <template #default="{ row }">
                        <el-tag :type="statusTagType(row.status)" effect="dark"> {{ statusText(row.status) }} </el-tag>
                    </template>
                </el-table-column>
                <el-table-column prop="notify_status" :label="t('common.notify_status')" :width="160">
                    <template #default="{ row }">
                        <el-tag :type="statusTagType(row.notify_status)" effect="dark"> {{ statusText(row.notify_status)
                        }} </el-tag>
                    </template>
                </el-table-column>
                <el-table-column :label="t('order.createdAt')" :width="180">
                    <template #default="{ row }">{{ formatTime(row.created_at) }}</template>
                </el-table-column>
                <el-table-column :label="t('common.actions')" :width="200" align="center">
                    <template #default="{ row }">
                        <el-button type="primary" link @click="showDetail(row)">{{ t('common.details') }}</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="pagination-wrapper">
                <el-pagination background :page-sizes="[10, 20, 50, 100]" :total="total" :pager-count="5"
                    v-model:current-page="pagination.index" v-model:page-size="pagination.size"
                    :layout="t('components.pagination.layout') || 'total, sizes, prev, pager, next, jumper'"
                    @current-change="loadData" @size-change="handleSizeChange" />
            </div>
        </el-card>
        <el-dialog v-model="detailVisible" :title="t('order.detailTitle')" width="800px" destroy-on-close>
            <el-descriptions :column="2" border>
                <el-descriptions-item :label="t('order.orderId')">{{ currentOrder?.id || '-' }}</el-descriptions-item>
                <el-descriptions-item :label="t('order.merchantId')">{{ currentOrder?.m_id || '-'
                    }}</el-descriptions-item>
                <el-descriptions-item :label="t('order.systemOrderNo')">{{ currentOrder?.no || '-'
                    }}</el-descriptions-item>
                <el-descriptions-item :label="t('order.merchantOrderNo')">{{ currentOrder?.c_no || '-'
                    }}</el-descriptions-item>
                <el-descriptions-item :label="t('order.amount')" :span="2"> {{ currentOrder?.amount }} {{
                    currentOrder?.coin }} </el-descriptions-item>
                <el-descriptions-item :label="t('order.currentStatus')" :span="2">
                    <el-tag :type="statusTagType(currentOrder?.status)"> {{ statusText(currentOrder?.status) }}
                    </el-tag>
                </el-descriptions-item>
                <el-descriptions-item :label="t('order.NotifyUrl')" :span="2">
                    <el-text type="info">{{ currentOrder?.notify_url || t('common.none') }}</el-text>
                </el-descriptions-item>
                <el-descriptions-item :label="t('order.createdAt')"> {{ formatTime(currentOrder?.created_at) }}
                </el-descriptions-item>
                <el-descriptions-item :label="t('order.updatedAt')"> {{ currentOrder?.updated_at ?
                    formatTime(currentOrder.updated_at) : '-' }} </el-descriptions-item>
                <el-descriptions-item :label="t('common.remark')" :span="2">
                    <el-text type="info">{{ currentOrder?.remark || t('common.noRemark') }}</el-text>
                </el-descriptions-item>
                <el-descriptions-item :label="t('order.deletedAt')" v-if="currentOrder?.deleted_at"> {{
                    formatTime(currentOrder.deleted_at) }} </el-descriptions-item>
            </el-descriptions>
            <template #footer>
                <el-button @click="detailVisible = false">{{ t('common.close') }}</el-button>
            </template>
        </el-dialog>
    </div>
</template>
<script setup>
import { ref, reactive, onMounted } from 'vue'
import api from '../../api'
import { ElMessage } from 'element-plus'
import dayjs from 'dayjs'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

// 新增商户选项
const merchantOptions = ref([])
const searchParams = reactive({
    keyword: '',
    merchantId: null, // 新增商户ID参数
    // ...其他已有参数
})

// 加载商户列表
const loadMerchants = async () => {
    try {
        const { data } = await api.get('/merchants/in')
        merchantOptions.value = data.map(m => ({
            value: m.id,
            label: m.nickname
        }))
    } finally {
    }
}

// 在原有代码基础上新增以下内容
const orderList = ref([])
const total = ref(0)
const loading = ref(false)
const pagination = reactive({
    index: 1,
    size: 10
})

// 状态显示处理
const statusTagType = (status) => {
    const map = {
        'success': t('status.success'),
        'pending': t('status.pending'),
        'failed': t('status.failed'),
        'canceled': t('status.canceled'),
        'listening': t('status.listening'),
        'timeout': t('status.timeout'),
        '': t('status.unknown')
    }
    return map[status] || 'info'
}

const statusText = (status) => {
    const map = {
        'success': t('statusText.completed'),
        'pending': t('statusText.created'),
        'failed': t('statusText.failed'),
        'canceled': t('statusText.canceled'),
        'listening': t('statusText.listening'),
        'timeout': t('statusText.timeout'),
        '': t('statusText.unknown')
    }
    return map[status] || t('statusText.unknown')
}

// 添加工具方法
const formatTime = (timeStr) => {
    return dayjs(timeStr).format('YYYY-MM-DD HH:mm:ss')
}

const getMerchantName = (mId) => {
    const merchant = merchantOptions.value.find(m => m.value === mId)
    return merchant?.label || '未知商户'
}

// 添加详情对话框数据
const detailVisible = ref(false)
const currentOrder = ref(null)

const showDetail = (row) => {
    currentOrder.value = row
    detailVisible.value = true
}

// 修改后的loadData方法
const loadData = async () => {
    try {
        loading.value = true
        const { data } = await api.post('/order/list', {
            page: pagination.index,
            size: pagination.size,
            order_no: searchParams.keyword,
            merchant_id: searchParams.merchantId
        })
        orderList.value = data.data
        total.value = data.total
    } catch (error) {
        ElMessage.error('加载订单失败' + error)
    } finally {
        loading.value = false
    }
}

const handleSizeChange = (newSize) => {
    pagination.size = newSize
    loadData()
}

onMounted(() => {
    loadData()
    loadMerchants() // 初始化时加载商户列表
})
</script>
<style scoped>
.filters {
    display: flex;
    align-items: center;
    gap: 10px;
}

.pagination-wrapper {
    margin-top: 20px;
    display: flex;
    justify-content: flex-end;
    padding: 10px 0;
    overflow-x: auto;
}

:deep(.el-pagination) {
    flex-wrap: wrap;
    gap: 8px;
}

:deep(.el-pagination__total) {
    margin-right: 0;
}

@media (max-width: 768px) {
    .pagination-wrapper {
        justify-content: center;
    }

    :deep(.el-pagination) {
        --el-pagination-button-width: 32px;
        --el-pagination-button-height: 32px;
    }

    :deep(.el-pagination__jump) {
        margin-left: 0;
        margin-top: 8px;
    }
}

.table-card {
    margin-top: 20px;
    border-radius: 8px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

:deep(.el-table) {
    font-size: 14px;

    th.el-table__cell {
        background-color: #f8f9fa;
        font-weight: 600;
    }

    .el-tag {
        font-size: 13px;
        padding: 0 8px;
    }
}

.toolbar-card {
    border-radius: 8px;
    margin-bottom: 20px;
}

:deep(.el-dialog) {
    max-width: 95vw;
}

:deep(.el-descriptions) {
    margin: 10px 0;
    font-size: 14px;
}

:deep(.el-descriptions__table) {
    width: 100%;
}

@media (max-width: 768px) {

    :deep(.el-descriptions__content) {
        flex: 1;
        min-width: calc(100% - 100px);
    }

    :deep(.el-dialog__footer) {
        text-align: center;
    }
}

@media (max-width: 768px) {
    :deep(.el-table td.el-table__cell) {
        padding: 8px 0;
    }

    :deep(.el-table .cell) {
        font-size: 13px;
    }

    :deep(.el-button--primary.is-link) {
        font-size: 12px;
        padding: 0 4px;
    }
}
</style>