export default {
    common: {
        logout: '退出登录',
        confirm: '确认',
        cancel: '取消',
        edit: '编辑',
        delete: '删除',
        create: '创建',
        search: '搜索',
        selectMerchant: '选择商户',
        close: '关闭',
        enabled: '启用',
        disabled: '停用',
        enable: '启用',
        disable: '停用',
        confirmTitle: '操作确认',
        confirmDelete: '确认删除',
        operationFailed: '操作失败',
        notSet: '未设置',
        home: '首页',
        profile: '个人中心',
        retry: '重试',
        confirmPayment: '确认支付 {amount} {coin}',
        backHome: '返回首页',
        language: '语言',
        zh: '中文',
        en: '英文',
        notify_status: '通知状态',
        selectMerchant: '选择商户'
    },
    login: {
        title: '欢迎回来 👋',
        subtitle: '请登录您的账户',
        email: '请输入邮箱',
        password: '请输入密码',
        submit: '立即登录'
    },
    dashboard: {
        dashboard: '仪表盘',
        merchantTotal: '总商户数',
        orderTotal: '总订单数',
        success7Days: '7天成功订单',
        totalAmount: '总交易额',
        welcome: '如果你喜欢cheems happy，请支持我们。pay.cheemshappy.com',
        platform: '支付管理平台'
    },
    order: {
        orders: '订单管理',
        searchOrder: '搜索订单号',
        orderNo: '订单号',
        merchantName: '商户名称',
        amount: '金额',
        createdAt: '创建时间',
        detailTitle: '订单详情',
        merchantId: '商户ID',
        systemOrderNo: '系统订单号',
        merchantOrderNo: '商户订单号',
        currentStatus: '当前状态',
        callbackUrl: '回调地址',
        notifyUrl: '通知地址',
        updatedAt: '最后更新时间',
        deletedAt: '删除时间',
        paymentOrder: '支付订单',
        merchant: '收款商户',
        targetNetwork: '目标网络',
        paymentSuccess: '支付成功！',
        autoRedirect: '秒后自动跳转至商户页面',
        jumpNow: '立即跳转'
    },
    status: {
        success: 'success',
        pending: 'warning',
        failed: 'danger',
        canceled: 'info',
        unknown: 'info',
        timeout: 'info'
    },
    statusText: {
        completed: '已完成',
        created: '已创建',
        failed: '失败',
        canceled: '已取消',
        unknown: '',
        timeout: '已超时'
    },
    components: {
        pagination: {
            layout: 'total, sizes, prev, pager, next, jumper',
            total: '共 {total} 条'
        }
    },
    merchants: {
        merchants: '商户管理',
        create: '新建商户',
        editTitle: '编辑商户',
        createTitle: '新建商户',
        merchantName: '商户名称',
        searchPlaceholder: '搜索商户名称',
        walletManagement: '钱包管理',
        rules: {
            nicknameRequired: '请输入商户名称',
            emailRequired: '请输入邮箱地址',
            emailFormat: '请输入正确的邮箱格式',
            phoneFormat: '请输入正确的手机号码'
        },
        updateSuccess: '商户信息更新成功',
        createSuccess: '商户创建成功',
        deleteSuccess: '商户删除成功',
        deleteFailed: '删除失败'
    },
    wallet: {
        wallets: '钱包管理',
        updateTitle: '更新钱包',
        createTitle: '创建钱包',
        address: '钱包地址',
        privateKey: '私钥',
        mnemonic: '助记词',
        autoGenerate: '自动生成',
        updateSuccess: '钱包更新成功',
        createSuccess: '钱包创建成功',
        autoGenerateFailed: '自动生成失败',
        create: '新增钱包',
        autoGenerateTitle: '自动生成钱包',
        selectMerchant: '选择商户',
        generateCount: '生成数量',
        startGenerate: '开始生成',
        generateSuccess: '成功生成{count}个钱包',
        generateFailed: '生成失败',
        statusUpdateSuccess: '状态更新成功',
        deleteConfirm: '确认删除该钱包？此操作不可逆！',
        disableConfirm: '确认停用该钱包？',
        enableConfirm: '确认启用该钱包？',
        merchant: '所属商户',
        balance: '余额',
        detailTitle: '钱包详情',
        connectMetaMask: '连接 MetaMask 钱包',
        connected: '已连接钱包',
        ensureNetwork: '请确保使用支持{network}的钱包'
    },
    settings: {
        systemWallet: '系统钱包设置',
        feeConfig: '费率配置',
        notification: '通知设置'
    },
    sysWallet: {
        mainAddress: '主地址',
        balance: '余额',
        lastSync: '最后同步',
        importWallet: '导入钱包',
        confirmImport: '确认导入',
        privateKey: '私钥',
        rules: {
            privateKeyRequired: '请输入私钥'
        },
        mnemonic: '助记词',
        currentIndex: '当前索引',
        createdAt: '创建时间',
        address: '地址',
        path: '路径',
        updateSuccess: '钱包更新成功',
        updateFailed: '更新失败：',
        importSuccess: '钱包导入成功',
        importFailed: '导入失败：',
        securityWarning: '私钥和助记词是系统钱包的重要信息，请妥善保管',
        longPressTip: '长按查看敏感信息，15秒后隐藏',
        noWallet: '未创建钱包',
        createWarning: '即将创建新的系统钱包，请注意：',
        createWarning1: '该操作将生成一个新的以太坊钱包',
        createWarning2: '请立即备份助记词和私钥',
        createWarning3: '该操作无法修改',
        createWarning4: '确保操作在安全环境下进行',
        walletManagement: '系统钱包管理'
    },
    otherSettings: {
        otherSettings: '其他系统设置',
        featureUnderDevelopment: '功能开发中'
    },
    systemSettings: {
        systemSettings: '系统设置',
        domain: '域名',
        domainPlaceholder: '请输入域名',
    },
    login: {
        login: '登录',
        welcome: '欢迎回来 👋',
        pleaseLogin: '请登录您的账户'
    },
    merchantsApi: {
        merchantsApi: '商户API管理',
        apikey: 'API密钥',
        callbackUrl: '回调地址',
        secretKey: '加密密钥',
        create: '新建API',
        editTitle: '编辑API配置',
        createTitle: '新建API配置',
        searchPlaceholder: '搜索商户名称/API密钥',
        deleteConfirm: '确认删除该API配置？此操作不可逆！',
        merchant: '所属商户'
    },
    rules: {
        required: '必填项',
        invalidDomain: '域名格式不正确'
    },
    error: {
        orderNotFound: '订单不存在',
        orderExpired: '该订单可能已完成支付或超过有效期',
        connectFailed: '钱包连接失败',
        paymentFailed: '支付失败'
    }
    // 其他翻译项...
} 