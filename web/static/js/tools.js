// web/static/js/tools.js
// 辅助工具管理页面 JavaScript
// 提供工具列表、配置、执行、输出查看等功能

// ============================================
// 全局状态
// ============================================

let toolsState = {
    tools: [],           // 工具列表
    categories: [],      // 分类列表
    selectedTool: null,  // 当前选中的工具
    executions: [],      // 执行记录
    runningExecution: null,  // 当前运行中的执行
    outputRefreshInterval: null  // 输出刷新定时器
};

// ============================================
// 页面初始化
// ============================================

// 页面加载完成后初始化
document.addEventListener('DOMContentLoaded', function() {
    // 检查是否切换到工具页面
    observePageSwitch();

    // 监听键盘快捷键
    document.addEventListener('keydown', handleKeyboardShortcuts);
});

// 观察页面切换
function observePageSwitch() {
    // 初始检查
    checkAndInitToolsPage();

    // 监听页面变化
    const observer = new MutationObserver(function(mutations) {
        mutations.forEach(function(mutation) {
            if (mutation.type === 'attributes' && mutation.attributeName === 'class') {
                const target = mutation.target;
                if (target.id && target.id.startsWith('page-')) {
                    if (target.classList.contains('active') && target.id === 'page-tools') {
                        initToolsPage();
                    }
                }
            }
        });
    });

    // 观察主内容区域
    const contentArea = document.querySelector('.content-area');
    if (contentArea) {
        observer.observe(contentArea, { attributes: true, subtree: true });
    }
}

// 检查并初始化工具页面
function checkAndInitToolsPage() {
    const toolsPage = document.getElementById('page-tools');
    if (toolsPage && toolsPage.classList.contains('active')) {
        initToolsPage();
    }
}

// ============================================
// 页面初始化
// ============================================

// initToolsPage 初始化工具页面
async function initToolsPage() {
    console.log('[Tools] 初始化辅助工具页面');

    // 加载工具列表
    await loadTools();

    // 加载执行历史
    await loadExecutions();

    // 设置事件监听
    setupEventListeners();
}

// loadTools 加载工具列表
async function loadTools() {
    try {
        const response = await apiFetch('/api/tools?enabled=true');
        if (!response.ok) throw new Error('获取工具列表失败');

        const data = await response.json();
        toolsState.tools = data.tools || [];

        // 渲染工具列表
        renderToolsList();

        // 渲染分类统计
        renderCategoryStats();

        console.log(`[Tools] 加载了 ${toolsState.tools.length} 个工具`);
    } catch (error) {
        console.error('[Tools] 加载工具失败:', error);
        showToast('加载工具列表失败: ' + error.message, 'error');
    }
}

// renderToolsList 渲染工具列表
function renderToolsList() {
    const container = document.getElementById('tools-grid');
    if (!container) return;

    if (toolsState.tools.length === 0) {
        container.innerHTML = `
            <div class="tools-empty">
                <p>暂无可用工具</p>
                <p class="tools-empty-hint">请在 tools 目录添加工具配置文件</p>
            </div>
        `;
        return;
    }

    // 按分类分组
    const grouped = {};
    toolsState.tools.forEach(tool => {
        const category = tool.category || '其他';
        if (!grouped[category]) {
            grouped[category] = [];
        }
        grouped[category].push(tool);
    });

    // 生成 HTML
    let html = '';
    Object.entries(grouped).forEach(([category, tools]) => {
        html += `
            <div class="tools-category">
                <h3 class="tools-category-title">
                    <span class="tools-category-icon">${getCategoryIcon(category)}</span>
                    ${escapeHtml(category)}
                    <span class="tools-category-count">(${tools.length})</span>
                </h3>
                <div class="tools-grid">
        `;

        tools.forEach(tool => {
            const isSelected = toolsState.selectedTool?.name === tool.name;
            html += `
                <div class="tool-card ${isSelected ? 'selected' : ''}"
                     data-tool-name="${escapeHtml(tool.name)}"
                     onclick="selectTool('${escapeHtml(tool.name)}')">
                    <div class="tool-card-header">
                        <span class="tool-name">${escapeHtml(tool.name)}</span>
                        <span class="tool-status-dot ${tool.enabled ? 'enabled' : 'disabled'}"></span>
                    </div>
                    <div class="tool-card-desc">${escapeHtml(tool.short_description || '')}</div>
                    <div class="tool-card-footer">
                        <span class="tool-params-count">
                            ${(tool.parameters || []).length} 个参数
                        </span>
                    </div>
                </div>
            `;
        });

        html += `
                </div>
            </div>
        `;
    });

    container.innerHTML = html;
}

// renderCategoryStats 渲染分类统计
function renderCategoryStats() {
    const container = document.getElementById('tools-stats');
    if (!container) return;

    const grouped = {};
    toolsState.tools.forEach(tool => {
        const category = tool.category || '其他';
        grouped[category] = (grouped[category] || 0) + 1;
    });

    let html = '<div class="tools-stats-row">';
    Object.entries(grouped).forEach(([category, count]) => {
        html += `
            <div class="tools-stat-item">
                <span class="tools-stat-icon">${getCategoryIcon(category)}</span>
                <span class="tools-stat-label">${escapeHtml(category)}</span>
                <span class="tools-stat-value">${count}</span>
            </div>
        `;
    });
    html += '</div>';

    container.innerHTML = html;
}

// ============================================
// 工具选择与参数配置
// ============================================

// selectTool 选择工具
async function selectTool(toolName) {
    const tool = toolsState.tools.find(t => t.name === toolName);
    if (!tool) return;

    toolsState.selectedTool = tool;

    // 更新工具卡片选中状态
    document.querySelectorAll('.tool-card').forEach(card => {
        card.classList.toggle('selected', card.dataset.toolName === toolName);
    });

    // 渲染工具详情面板
    renderToolPanel(tool);

    console.log(`[Tools] 选中工具: ${toolName}`);
}

// renderToolPanel 渲染工具配置面板
function renderToolPanel(tool) {
    const container = document.getElementById('tool-panel');
    if (!container) return;

    // 构建参数表单
    let paramsHtml = '';
    (tool.parameters || []).forEach(param => {
        const isRequired = param.required ? '<span class="required">*</span>' : '';
        const defaultValue = param.default !== undefined ? param.default : '';

        if (param.type === 'bool') {
            paramsHtml += `
                <div class="form-group">
                    <label class="checkbox-label">
                        <input type="checkbox" id="param-${param.name}"
                               class="modern-checkbox" ${defaultValue ? 'checked' : ''}>
                        <span class="checkbox-custom"></span>
                        <span>${escapeHtml(param.name)}</span>
                    </label>
                    <small class="form-hint">${escapeHtml(param.description || '')}</small>
                </div>
            `;
        } else if (param.format === 'flag') {
            // flag 格式单独一行，label是参数名
            paramsHtml += `
                <div class="form-group">
                    <label for="param-${param.name}">
                        ${escapeHtml(param.name)} ${isRequired}
                        <code class="param-flag">${escapeHtml(param.flag || '')}</code>
                    </label>
                    <input type="${param.type === 'int' ? 'number' : 'text'}"
                           id="param-${param.name}"
                           class="form-control"
                           placeholder="${escapeHtml(param.description || '')}"
                           value="${escapeHtml(String(defaultValue))}">
                    <small class="form-hint">${escapeHtml(param.description || '')}</small>
                </div>
            `;
        } else {
            // positional/其他格式
            paramsHtml += `
                <div class="form-group">
                    <label for="param-${param.name}">
                        ${escapeHtml(param.name)} ${isRequired}
                    </label>
                    <input type="${param.type === 'int' ? 'number' : 'text'}"
                           id="param-${param.name}"
                           class="form-control"
                           placeholder="${escapeHtml(param.description || '')}"
                           value="${escapeHtml(String(defaultValue))}">
                    <small class="form-hint">${escapeHtml(param.description || '')}</small>
                </div>
            `;
        }
    });

    // 添加强制终止复选框（用于危险命令）
    const isDangerous = isDangerousTool(tool.name);
    const dangerWarning = isDangerous ? `
        <div class="tool-danger-warning">
            <span class="warning-icon">⚠️</span>
            <span>这是一个危险工具，执行前请确认目标授权</span>
        </div>
    ` : '';

    container.innerHTML = `
        <div class="tool-panel-header">
            <h3>${escapeHtml(tool.name)}</h3>
            <button class="btn-icon" onclick="closeToolPanel()" title="关闭">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <line x1="18" y1="6" x2="6" y2="18"></line>
                    <line x1="6" y1="6" x2="18" y2="18"></line>
                </svg>
            </button>
        </div>

        <div class="tool-panel-desc">
            <p>${escapeHtml(tool.description || tool.short_description || '')}</p>
        </div>

        ${dangerWarning}

        <div class="tool-panel-section">
            <h4>命令预览</h4>
            <div class="command-preview">
                <code id="command-preview-text">${escapeHtml(tool.command)} ${(tool.args || []).join(' ')}</code>
                <button class="btn-icon" onclick="copyCommand()" title="复制命令">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
                        <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
                    </svg>
                </button>
            </div>
        </div>

        <div class="tool-panel-section">
            <h4>参数配置</h4>
            <div class="tool-params-form" id="tool-params-form">
                ${paramsHtml || '<p class="no-params">此工具无需额外参数</p>'}
            </div>
        </div>

        <div class="tool-panel-section">
            <h4>执行选项</h4>
            <div class="form-group">
                <label for="tool-timeout">超时时间（秒）</label>
                <input type="number" id="tool-timeout" class="form-control" value="300" min="0" max="3600">
                <small class="form-hint">设置为 0 表示使用默认超时 (300秒)</small>
            </div>
            <div class="form-group">
                <label for="tool-workdir">工作目录</label>
                <input type="text" id="tool-workdir" class="form-control" placeholder="可选，留空使用默认目录">
            </div>
            <div class="form-group">
                <label for="tool-extra-args">额外参数</label>
                <input type="text" id="tool-extra-args" class="form-control"
                       placeholder="直接追加到命令末尾">
                <small class="form-hint">直接追加到命令末尾，不经过参数解析</small>
            </div>
        </div>

        <div class="tool-panel-actions">
            <button class="btn-secondary" onclick="clearToolForm()">清空</button>
            <button class="btn-primary" onclick="executeSelectedTool()">
                <span class="btn-icon-run">▶</span> 执行
            </button>
        </div>
    `;

    // 为参数输入添加实时预览更新
    setupParamListeners(tool);
}

// setupParamListeners 设置参数监听器，实时更新命令预览
function setupParamListeners(tool) {
    const form = document.getElementById('tool-params-form');
    if (!form) return;

    // 监听所有输入变化
    form.addEventListener('input', () => updateCommandPreview(tool));
    form.addEventListener('change', () => updateCommandPreview(tool));

    // 监听额外参数
    const extraArgs = document.getElementById('tool-extra-args');
    if (extraArgs) {
        extraArgs.addEventListener('input', () => updateCommandPreview(tool));
    }
}

// updateCommandPreview 更新命令预览
function updateCommandPreview(tool) {
    const preview = document.getElementById('command-preview-text');
    if (!preview) return;

    // 构建参数
    const args = buildArgsFromForm(tool);
    const extraArgs = document.getElementById('tool-extra-args');
    const extra = extraArgs?.value?.trim() || '';

    let cmd = `${tool.command} ${args.join(' ')}`;
    if (extra) {
        cmd += ` ${extra}`;
    }

    preview.textContent = cmd;
}

// buildArgsFromForm 从表单构建参数
function buildArgsFromForm(tool) {
    const args = [...(tool.args || [])];

    (tool.parameters || []).forEach(param => {
        const input = document.getElementById(`param-${param.name}`);
        if (!input) return;

        let value;
        if (param.type === 'bool') {
            value = input.checked;
        } else {
            value = input.value?.trim();
        }

        // 跳过空值
        if (value === '' || value === undefined) {
            // 使用默认值
            if (param.default !== undefined) {
                value = param.default;
            } else {
                return;
            }
        }

        if (param.type === 'bool') {
            if (value) {
                if (param.flag) {
                    args.push(param.flag);
                }
            }
        } else if (param.format === 'flag' && param.flag) {
            args.push(param.flag, String(value));
        } else if (param.format === 'combined' && param.flag) {
            args.push(`${param.flag}=${value}`);
        } else {
            args.push(String(value));
        }
    });

    return args;
}

// copyCommand 复制命令
function copyCommand() {
    const preview = document.getElementById('command-preview-text');
    if (!preview) return;

    navigator.clipboard.writeText(preview.textContent).then(() => {
        showToast('命令已复制到剪贴板', 'success');
    }).catch(err => {
        console.error('复制失败:', err);
    });
}

// closeToolPanel 关闭工具面板
function closeToolPanel() {
    const container = document.getElementById('tool-panel');
    if (container) {
        container.innerHTML = `
            <div class="tool-panel-empty">
                <div class="empty-icon">🔧</div>
                <p>选择一个工具开始</p>
                <p class="hint">点击左侧工具卡片查看详情并配置参数</p>
            </div>
        `;
    }
    toolsState.selectedTool = null;

    // 移除选中状态
    document.querySelectorAll('.tool-card').forEach(card => {
        card.classList.remove('selected');
    });
}

// clearToolForm 清空表单
function clearToolForm() {
    const form = document.getElementById('tool-params-form');
    if (!form) return;

    // 清空所有输入
    form.querySelectorAll('input').forEach(input => {
        if (input.type === 'checkbox') {
            input.checked = false;
        } else {
            input.value = '';
        }
    });

    // 清空额外参数
    const extraArgs = document.getElementById('tool-extra-args');
    if (extraArgs) extraArgs.value = '';

    // 重置超时和工作目录
    const timeout = document.getElementById('tool-timeout');
    if (timeout) timeout.value = '300';

    const workdir = document.getElementById('tool-workdir');
    if (workdir) workdir.value = '';

    // 更新命令预览
    if (toolsState.selectedTool) {
        updateCommandPreview(toolsState.selectedTool);
    }
}

// ============================================
// 工具执行
// ============================================

// executeSelectedTool 执行选中的工具
async function executeSelectedTool() {
    const tool = toolsState.selectedTool;
    if (!tool) {
        showToast('请先选择一个工具', 'warning');
        return;
    }

    // 危险工具二次确认
    if (isDangerousTool(tool.name)) {
        const confirmed = confirm('这是一个危险工具，确定要执行吗？\n\n工具: ' + tool.name + '\n请确保您有权对此目标进行测试。');
        if (!confirmed) return;
    }

    // 构建参数
    const args = buildArgsFromForm(tool);
    const timeout = parseInt(document.getElementById('tool-timeout')?.value || '300');
    const workdir = document.getElementById('tool-workdir')?.value?.trim() || '';

    // 显示执行中的状态
    const actionsContainer = document.querySelector('.tool-panel-actions');
    if (actionsContainer) {
        actionsContainer.innerHTML = `
            <button class="btn-secondary" disabled>清空</button>
            <button class="btn-danger" onclick="cancelCurrentExecution()">
                <span class="btn-icon-stop">⏹</span> 取消
            </button>
            <span class="running-indicator">执行中...</span>
        `;
    }

    try {
        const response = await apiFetch(`/api/tools/${tool.name}/execute`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                args: argsToObject(tool, args),
                timeout: timeout,
                working_dir: workdir
            })
        });

        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || '执行失败');
        }

        const result = await response.json();
        toolsState.runningExecution = result.execution_id;

        console.log(`[Tools] 提交执行: ${result.execution_id}`);

        // 开始轮询执行状态
        startOutputPolling(result.execution_id);

        // 添加到历史
        addToHistory(result.execution_id, tool.name, args);

    } catch (error) {
        console.error('[Tools] 执行失败:', error);
        showToast('执行失败: ' + error.message, 'error');

        // 恢复按钮
        restoreExecuteButton();
    }
}

// argsToObject 将参数数组转换为对象
function argsToObject(tool, args) {
    const obj = {};

    (tool.parameters || []).forEach((param, index) => {
        if (param.position && param.position <= args.length) {
            obj[param.name] = args[param.position - 1];
        }
    });

    return obj;
}

// startOutputPolling 开始轮询输出
function startOutputPolling(execId) {
    // 先获取初始状态
    pollExecution(execId);

    // 设置定时器
    if (toolsState.outputRefreshInterval) {
        clearInterval(toolsState.outputRefreshInterval);
    }

    toolsState.outputRefreshInterval = setInterval(() => {
        pollExecution(execId);
    }, 1000);
}

// pollExecution 轮询执行状态
async function pollExecution(execId) {
    try {
        const response = await apiFetch(`/api/tools/executions/${execId}`);
        if (!response.ok) return;

        const data = await response.json();
        const exec = data.execution;

        // 更新输出显示
        renderExecutionOutput(exec);

        // 检查是否完成
        if (exec.status !== 'running' && exec.status !== 'pending') {
            // 执行结束
            clearInterval(toolsState.outputRefreshInterval);
            toolsState.outputRefreshInterval = null;
            toolsState.runningExecution = null;

            // 刷新历史
            loadExecutions();

            // 恢复按钮
            restoreExecuteButton();

            // 显示完成提示
            if (exec.status === 'completed') {
                showToast(`工具执行完成，退出码: ${exec.exit_code}`, 'success');
            } else if (exec.status === 'failed') {
                showToast('工具执行失败: ' + (exec.error || '未知错误'), 'error');
            } else if (exec.status === 'cancelled') {
                showToast('工具执行已取消', 'warning');
            }
        }
    } catch (error) {
        console.error('[Tools] 轮询失败:', error);
    }
}

// renderExecutionOutput 渲染执行输出
function renderExecutionOutput(exec) {
    const container = document.getElementById('tool-output');
    if (!container) return;

    const stdout = exec.stdout || '';
    const stderr = exec.stderr || '';

    container.innerHTML = `
        <div class="output-header">
            <span class="output-status ${exec.status}">${getStatusText(exec.status)}</span>
            <span class="output-info">
                ${exec.exit_code !== undefined ? `<span class="exit-code">退出码: ${exec.exit_code}</span>` : ''}
                ${exec.duration_ms ? `<span class="duration">时长: ${(exec.duration_ms / 1000).toFixed(2)}s</span>` : ''}
            </span>
        </div>
        <div class="output-content">
            <div class="output-tabs">
                <button class="output-tab active" onclick="showOutputTab('stdout')">标准输出</button>
                <button class="output-tab" onclick="showOutputTab('stderr')">错误输出</button>
                <button class="output-tab" onclick="showOutputTab('all')">全部</button>
            </div>
            <div class="output-body">
                <pre id="output-text" class="output-pre">${exec.status === 'running' || exec.status === 'pending' ? '等待输出...' : escapeHtml(stdout || '无标准输出')}</pre>
                <pre id="output-stderr" class="output-pre hidden">${escapeHtml(stderr || '无错误输出')}</pre>
            </div>
        </div>
        ${exec.status === 'running' || exec.status === 'pending' ? `
            <div class="output-loading">
                <div class="loading-spinner"></div>
                <span>工具运行中...</span>
            </div>
        ` : ''}
    `;
}

// showOutputTab 显示输出标签页
function showOutputTab(tab) {
    const stdout = document.getElementById('output-text');
    const stderr = document.getElementById('output-stderr');

    document.querySelectorAll('.output-tab').forEach(t => t.classList.remove('active'));
    event.target.classList.add('active');

    if (tab === 'stdout') {
        stdout?.classList.remove('hidden');
        stderr?.classList.add('hidden');
    } else if (tab === 'stderr') {
        stdout?.classList.add('hidden');
        stderr?.classList.remove('hidden');
    } else {
        stdout?.classList.remove('hidden');
        stderr?.classList.remove('hidden');
    }
}

// restoreExecuteButton 恢复执行按钮
function restoreExecuteButton() {
    const actionsContainer = document.querySelector('.tool-panel-actions');
    if (actionsContainer && toolsState.selectedTool) {
        const isDangerous = isDangerousTool(toolsState.selectedTool.name);
        actionsContainer.innerHTML = `
            <button class="btn-secondary" onclick="clearToolForm()">清空</button>
            <button class="btn-primary" onclick="executeSelectedTool()">
                <span class="btn-icon-run">▶</span> 执行
            </button>
        `;
    }
}

// cancelCurrentExecution 取消当前执行
async function cancelCurrentExecution() {
    if (!toolsState.runningExecution) return;

    try {
        const response = await apiFetch(`/api/tools/executions/${toolsState.runningExecution}/cancel`, {
            method: 'POST'
        });

        if (response.ok) {
            showToast('已发送取消请求', 'success');
        }
    } catch (error) {
        console.error('[Tools] 取消失败:', error);
    }
}

// ============================================
// 执行历史
// ============================================

// loadExecutions 加载执行历史
async function loadExecutions() {
    try {
        const response = await apiFetch('/api/tools/executions?page=1&page_size=50');
        if (!response.ok) throw new Error('获取执行历史失败');

        const data = await response.json();
        toolsState.executions = data.executions || [];

        renderExecutionsHistory();
    } catch (error) {
        console.error('[Tools] 加载历史失败:', error);
    }
}

// renderExecutionsHistory 渲染执行历史
function renderExecutionsHistory() {
    const container = document.getElementById('executions-list');
    if (!container) return;

    if (toolsState.executions.length === 0) {
        container.innerHTML = `
            <div class="history-empty">
                <p>暂无执行记录</p>
            </div>
        `;
        return;
    }

    let html = '';
    toolsState.executions.forEach(exec => {
        const duration = exec.duration_ms ? `${(exec.duration_ms / 1000).toFixed(1)}s` : '-';
        html += `
            <div class="execution-item ${exec.status}" onclick="viewExecution('${exec.id}')">
                <div class="execution-item-header">
                    <span class="execution-tool">${escapeHtml(exec.tool_name)}</span>
                    <span class="execution-status ${exec.status}">${getStatusText(exec.status)}</span>
                </div>
                <div class="execution-item-info">
                    <span class="execution-id">${exec.id}</span>
                    <span class="execution-time">${formatTime(exec.start_time)}</span>
                    <span class="execution-duration">${duration}</span>
                </div>
                ${exec.exit_code !== undefined ? `
                    <div class="execution-item-exitcode ${exec.exit_code === 0 ? 'success' : 'failed'}">
                        退出码: ${exec.exit_code}
                    </div>
                ` : ''}
            </div>
        `;
    });

    container.innerHTML = html;
}

// viewExecution 查看执行详情
async function viewExecution(execId) {
    try {
        const response = await apiFetch(`/api/tools/executions/${execId}`);
        if (!response.ok) throw new Error('获取执行详情失败');

        const data = await response.json();
        renderExecutionOutput(data.execution);

        // 选中输出面板
        const outputPanel = document.getElementById('tool-output');
        if (outputPanel) {
            outputPanel.scrollIntoView({ behavior: 'smooth' });
        }
    } catch (error) {
        console.error('[Tools] 查看执行详情失败:', error);
    }
}

// addToHistory 添加到历史列表
function addToHistory(execId, toolName, args) {
    const historyItem = {
        id: execId,
        tool_name: toolName,
        args: args,
        start_time: new Date().toISOString(),
        status: 'pending'
    };

    toolsState.executions.unshift(historyItem);

    // 限制历史数量
    if (toolsState.executions.length > 100) {
        toolsState.executions = toolsState.executions.slice(0, 100);
    }

    renderExecutionsHistory();
}

// ============================================
// 辅助函数
// ============================================

// getCategoryIcon 获取分类图标
function getCategoryIcon(category) {
    const icons = {
        '信息收集': '🔍',
        '漏洞检测': '🛡️',
        '暴力破解': '🔓',
        '密码攻击': '🔑',
        'Web渗透': '🌐',
        '逆向工程': '🔬',
        '云安全': '☁️',
        '容器安全': '📦',
        '其他': '🔧'
    };
    return icons[category] || '🔧';
}

// getStatusText 获取状态文本
function getStatusText(status) {
    const texts = {
        'pending': '等待中',
        'running': '运行中',
        'completed': '已完成',
        'failed': '失败',
        'cancelled': '已取消'
    };
    return texts[status] || status;
}

// isDangerousTool 检查是否为危险工具
function isDangerousTool(toolName) {
    const dangerous = ['hydra', 'hashcat', 'john', 'msfvenom', 'metasploit', 'netexec'];
    return dangerous.includes(toolName.toLowerCase());
}

// formatTime 格式化时间
function formatTime(isoString) {
    if (!isoString) return '-';
    const date = new Date(isoString);
    return date.toLocaleString('zh-CN', {
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit'
    });
}

// escapeHtml HTML转义
function escapeHtml(str) {
    if (!str) return '';
    const div = document.createElement('div');
    div.textContent = str;
    return div.innerHTML;
}

// showToast 显示提示消息
function showToast(message, type) {
    // 使用现有的 toast 机制或创建简单的 toast
    const toast = document.createElement('div');
    toast.className = `toast toast-${type}`;
    toast.textContent = message;

    let container = document.querySelector('.toast-container');
    if (!container) {
        container = document.createElement('div');
        container.className = 'toast-container';
        document.body.appendChild(container);
    }

    container.appendChild(toast);

    setTimeout(() => {
        toast.classList.add('show');
    }, 10);

    setTimeout(() => {
        toast.classList.remove('show');
        setTimeout(() => toast.remove(), 300);
    }, 3000);
}

// setupEventListeners 设置事件监听
function setupEventListeners() {
    // 搜索框
    const searchInput = document.getElementById('tools-search');
    if (searchInput) {
        searchInput.addEventListener('input', debounce(filterTools, 300));
    }

    // 分类筛选
    const categoryFilter = document.getElementById('tools-category-filter');
    if (categoryFilter) {
        categoryFilter.addEventListener('change', filterTools);
    }
}

// filterTools 过滤工具
function filterTools() {
    const search = document.getElementById('tools-search')?.value?.toLowerCase() || '';
    const category = document.getElementById('tools-category-filter')?.value || '';

    document.querySelectorAll('.tool-card').forEach(card => {
        const name = card.dataset.toolName?.toLowerCase() || '';
        const cardCategory = card.closest('.tools-category')?.querySelector('h3')?.textContent || '';

        const matchSearch = !search || name.includes(search);
        const matchCategory = !category || cardCategory.includes(category);

        card.style.display = matchSearch && matchCategory ? '' : 'none';
    });
}

// debounce 防抖
function debounce(func, wait) {
    let timeout;
    return function executedFunction(...args) {
        const later = () => {
            clearTimeout(timeout);
            func(...args);
        };
        clearTimeout(timeout);
        timeout = setTimeout(later, wait);
    };
}

// handleKeyboardShortcuts 处理键盘快捷键
function handleKeyboardShortcuts(e) {
    // Ctrl/Cmd + Enter 执行当前工具
    if ((e.ctrlKey || e.metaKey) && e.key === 'Enter') {
        const toolsPage = document.getElementById('page-tools');
        if (toolsPage && toolsPage.classList.contains('active')) {
            e.preventDefault();
            executeSelectedTool();
        }
    }

    // Escape 关闭面板
    if (e.key === 'Escape') {
        closeToolPanel();
    }
}
