{{ define "actions" }}
<div class="row col-md-9">
 <form method="POST">
  <div class="form-group">
   <button type="submit" class="btn btn-default btn-xs" name="edit" value="true">
    <span class="glyphicon glyphicon-edit"></span>Изменить
   </button>
   <button type="submit" class="btn btn-default btn-xs" name="revisions" value="true">
    <span class="glyphicon glyphicon-list-alt"></span> Изменения 
   </button>
  </div>
 </form>
</div>
{{end}}
