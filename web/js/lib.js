var myApp = angular.module('recordsApp', ['ui.bootstrap']);

myApp.controller('recordsController', function($scope, $http) {
	$scope.filteredRecords = []
	,$scope.currentPage = 1
	,$scope.numPerPage = 10
	,$scope.maxSize = 5;

	$scope.pagination = function() {
		console.log($scope.records.length)
		$scope.$watch("currentPage + numPerPage", function() {
	    	var begin = (($scope.currentPage - 1) * $scope.numPerPage), end = begin + $scope.numPerPage;
		    $scope.filteredRecords = $scope.records.slice(begin, end);
		});
	};

	$scope.getRecords = function() {
		$scope.loading = true;
		var path = "/records"
		var date = document.getElementsByClassName("calendar")[0].value;

		if ( date !== "" ) {
			path = path + "?date=" + date;
		} 

		$http.get(path).then(function successCallback(response) {
	  		$scope.records = response.data;

			$scope.$watch("currentPage + numPerPage", function() {
		    	var begin = (($scope.currentPage - 1) * $scope.numPerPage), end = begin + $scope.numPerPage;
			    $scope.filteredRecords = $scope.records.slice(begin, end);
			});
			$scope.loading = false;
	  	});
	};

	$scope.getRecords();







	pickmeup.defaults.locales['ru'] = {
		days: ['Воскресенье', 'Понедельник', 'Вторник', 'Среда', 'Четверг', 'Пятница', 'Суббота'],
		daysShort: ['Вс', 'Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб'],
		daysMin: ['Вс', 'Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб'],
		months: ['Январь', 'Февраль', 'Март', 'Апрель', 'Май', 'Июнь', 'Июль', 'Август', 'Сентябрь', 'Октябрь', 'Ноябрь', 'Декабрь'],
		monthsShort: ['Янв', 'Фев', 'Мар', 'Апр', 'Май', 'Июн', 'Июл', 'Авг', 'Сен', 'Окт', 'Ноя', 'Дек']
	};

	pickmeup.defaults.locale = "ru";

	var calendar = document.getElementsByClassName("calendar")[0];

	pickmeup(calendar, {
	  position       : 'right',
	  hide_on_select : true,
	  format	: 'Y-m-d'
	});

	calendar.addEventListener('pickmeup-change', function (e) {
	    $scope.getRecords();
	})

});


