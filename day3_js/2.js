/*����һ�����Ͳ��� n�������д������һ�� counter ������
��� counter ����������� n��
ÿ�ε�����ʱ�᷵��ǰһ��ֵ�� 1 ��ֵ ( n ,  n + 1 ,  n + 2 ���ȵ�)��*/

var createCounter = function(n) {
    
    return function() {
        return n++;

    };
};