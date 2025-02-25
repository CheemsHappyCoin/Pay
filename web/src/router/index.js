import { createRouter, createWebHistory, createWebHashHistory } from 'vue-router'
import Login from '../pages/Login.vue'
import AdminLayout from '../layouts/AdminLayout.vue'
import Dashboard from '../pages/admin/Dashboard.vue'

const routes = [
    {
        path: '/login',
        name: 'Login',
        component: Login
    },
    {
        path: '/admin',
        component: AdminLayout,
        meta: { requiresAuth: true },
        children: [
            {
                path: 'dashboard',
                name: 'Dashboard',
                component: Dashboard
            },
            {
                path: 'merchants',
                name: 'MerchantsManagement',
                component: () => import('../pages/admin/MerchantsManagement.vue'),
                meta: { requiresAuth: true }
            },
            {
                path: 'wallets',
                name: 'WalletManagement',
                component: () => import('../pages/admin/WalletManagement.vue'),
                meta: { requiresAuth: true }
            },
            {
                path: 'system-settings',
                component: () => import('../pages/admin/SystemSettings.vue'),
                redirect: '/admin/system-settings/sys-wallet',
                children: [
                    {
                        path: 'sys-wallet',
                        name: 'sysWallet',
                        component: () => import('../pages/admin/settings/SysWalletSettings.vue')
                    },
                    {
                        path: 'user-settings',
                        name: 'userSettings',
                        component: () => import('../pages/admin/settings/UserSettings.vue')
                    },
                    {
                        path: 'other-settings',
                        name: 'otherSettings',
                        component: () => import('../pages/admin/settings/OtherSettings.vue')
                    }
                ]
            },
            {
                path: 'orders',
                name: 'OrdersManagement',
                component: () => import('../pages/admin/OrderManagement.vue'),
                meta: { requiresAuth: true }
            },
            {
                path: 'merchants-api',
                name: 'MerchantsApiManagement',
                component: () => import('../pages/admin/MerchantsApiManagement.vue'),
                meta: { requiresAuth: true }
            },
        ]
    },
    {
        path: '/cheems/happy/pay',
        name: 'PayPage',
        component: () => import('../pages/PayPage.vue'),
        meta: { requiresAuth: false }
    },
    {
        path: '/',
        redirect: '/admin/dashboard'
    }
]

const router = createRouter({
    history: import.meta.env.VITE_ROUTER_MODE === 'hash'
        ? createWebHashHistory()
        : createWebHistory(),
    routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
    const isAuthenticated = localStorage.getItem('token')
    if (to.meta.requiresAuth && !isAuthenticated) {
        next('/login')
    } else {
        next()
    }
})

export default router 