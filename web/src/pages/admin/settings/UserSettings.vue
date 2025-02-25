<template>
    <el-card class="user-settings">
        <template #header>
            <div class="card-header">
                <h2>🔐 {{ t('userSettings.userSettings') }}</h2>
            </div>
        </template>
        <el-form :model="form" :rules="rules" ref="formRef" label-width="120px" label-position="left"
            @submit.prevent="handleSubmit">
            <!-- 邮箱/账号 -->
            <el-form-item :label="t('userSettings.email')" prop="email">
                <el-input v-model="form.email" :placeholder="t('userSettings.emailPlaceholder')" clearable type="email"
                    :disabled="!isEditing">
                    <template #prefix>
                        <i class="el-icon-message"></i>
                    </template>
                </el-input>
            </el-form-item>
            <!-- 当前密码 -->
            <el-form-item v-if="isEditing" :label="t('userSettings.currentPassword')" prop="currentPassword">
                <el-input v-model="form.currentPassword" :placeholder="t('userSettings.currentPasswordPlaceholder')"
                    show-password clearable>
                </el-input>
            </el-form-item>
            <!-- 新密码 -->
            <el-form-item v-if="isEditing" :label="t('userSettings.newPassword')" prop="newPassword">
                <el-input v-model="form.newPassword" :placeholder="t('userSettings.newPasswordPlaceholder')"
                    show-password clearable>
                </el-input>
            </el-form-item>
            <el-form-item>
                <el-button v-if="!isEditing" type="primary" @click="isEditing = true"> {{ t('common.edit') }}
                </el-button>
                <el-button v-else type="primary" native-type="submit" :loading="submitting"> {{ t('common.save') }}
                </el-button>
                <el-button v-if="isEditing" @click="cancelEdit"> {{ t('common.cancel') }} </el-button>
            </el-form-item>
        </el-form>
    </el-card>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useI18n } from 'vue-i18n'
import api from '../../../api'

const { t } = useI18n()
const form = ref({
    email: '',
    currentPassword: '',
    newPassword: ''
})
const isEditing = ref(false)
const formRef = ref(null)
const submitting = ref(false)

// 获取用户信息
const fetchProfile = async () => {
    try {
        const { data } = await api.get('/user/profile')
        console.log(data)
        form.value.email = data.email
    } catch (error) {
        ElMessage.error(error.response?.data?.message || '获取用户信息失败')
    }
}

// 表单验证规则
const rules = {
    email: [
        { required: true, message: t('userSettings.emailRequired'), trigger: 'blur' },
        { type: 'email', message: t('userSettings.emailInvalid'), trigger: ['blur', 'change'] }
    ],
    currentPassword: [
        { required: true, message: t('userSettings.currentPasswordRequired'), trigger: 'blur' }
    ],
    newPassword: [
        {
            validator: (_, value, callback) => {
                if (value && value.length < 8) {
                    callback(new Error(t('userSettings.passwordLength')))
                } else {
                    callback()
                }
            },
            trigger: ['blur', 'change']
        }
    ]
}

// 提交修改
const handleSubmit = async () => {
    try {
        await formRef.value.validate()
        submitting.value = true

        const payload = {
            email: form.value.email,
            password: form.value.currentPassword,
            re_password: form.value.newPassword
        }

        await api.post('/user/update', payload)

        ElMessage.success(t('userSettings.updateSuccess'))
        isEditing.value = false
        form.value.currentPassword = ''
        form.value.newPassword = ''
    } catch (error) {
        const msg = error.response?.data?.message || t('userSettings.updateFailed')
        ElMessage.error(msg)
    } finally {
        submitting.value = false
    }
}

// 取消编辑
const cancelEdit = () => {
    isEditing.value = false
    fetchProfile() // 重置表单数据
    form.value.currentPassword = ''
    form.value.newPassword = ''
}

// 初始化获取用户信息
onMounted(fetchProfile)
</script>
<style scoped>
.user-settings {
    max-width: 600px;
    margin: 20px auto;
    padding: 20px;
}

.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.el-form-item {
    margin-bottom: 22px;
}
</style>