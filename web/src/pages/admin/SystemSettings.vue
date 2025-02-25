<template>
    <div class="system-settings">
        <el-card class="nav-card">
            <div class="nav-menu">
                <el-menu :default-active="activeTab" class="vertical-menu" router @select="handleMenuSelect">
                    <el-menu-item index="/admin/system-settings/sys-wallet">
                        <i class="el-icon-wallet"></i>
                        <span>{{ t('sysWallet.walletManagement') }}</span>
                    </el-menu-item>
                    <el-menu-item index="/admin/system-settings/other-settings">
                        <i class="el-icon-setting"></i>
                        <span>{{ t('otherSettings.otherSettings') }}</span>
                    </el-menu-item>
                    <el-menu-item index="/admin/system-settings/user-settings">
                        <i class="el-icon-setting"></i>
                        <span>{{ t('userSettings.userSettings') }}</span>
                    </el-menu-item>
                </el-menu>
            </div>
        </el-card>
        <div class="content-area">
            <router-view style="min-height: 400px" />
        </div>
    </div>
</template>
<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'

const route = useRoute()
const activeTab = ref('')
const { t } = useI18n()

watch(() => route.path, (newPath) => {
    activeTab.value = newPath
}, { immediate: true })

const handleMenuSelect = (index) => {
    // This function is no longer used in the new template
}
</script>
<style scoped lang="scss">
.system-settings {
    display: flex;
    gap: 20px;

    .nav-card {
        width: 240px;

        :deep(.el-card__body) {
            padding: 0;
        }
    }

    .content-area {
        flex: 1;
        min-width: 0;
    }

    .vertical-menu {
        border-right: none;

        .el-menu-item {
            height: 50px;
            line-height: 50px;
            font-size: 14px;

            i {
                margin-right: 8px;
            }
        }
    }
}
</style>